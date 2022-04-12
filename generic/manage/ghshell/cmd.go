package ghshell

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/pkg/errors"
)

const defaultShell = "zsh"

var (
	whichsh string
	shFMT   string
)

func init() {
	whichsh = os.ExpandEnv(`$SHELL`)
	if whichsh == "" {
		whichsh = defaultShell
	}
	shFMT = whichsh + " %s"
}

type (
	Sheller interface {
		Out(args ...string) (result string, err error)
		Combined(args ...string) (result string, err error)
		Quick(args ...string) string
	}

	Cmd struct {
		ctx    context.Context
		cancel context.CancelFunc

		stdcmd *exec.Cmd

		sync.Mutex

		Bash      string
		ShellMode bool
		Status    Status
		Env       []string
		Dir       string

		isFinalized bool

		timeout int

		statusChan chan Status
		doneChan   chan error

		output bytes.Buffer // stdout + stderr
		stdout bytes.Buffer
		stderr bytes.Buffer
	}

	Status struct {
		PID      int
		Finish   bool
		ExitCode int
		Error    error
		CostTime time.Duration

		Output string // stdout + stderr
		Stdout string
		Stderr string

		startTime time.Time
		endTime   time.Time
	}

	result struct {
		stdout string
		stderr string
		errNo  int
		err    error
		done   chan bool
	}

	optionFunc func(*Cmd) error
)

func NewCommand(bash string, options ...optionFunc) *Cmd {
	c := &Cmd{
		Bash:       bash,
		Status:     Status{},
		ShellMode:  true,
		statusChan: make(chan Status, 1),
		doneChan:   make(chan error, 1),
	}
	for _, opt := range options {
		opt(c)
	}
	return c
}

// Start async execute command
func (c *Cmd) Start() error {
	if c.Status.Finish {
		return ErrAlreadyFinished
	}

	return c.run()
}

// Wait waits for command to finish
func (c *Cmd) Wait() error {
	<-c.doneChan
	return c.Status.Error
}

// Run starts and waits for process exit
func (c *Cmd) Run() error {
	c.Start()
	return c.Wait()
}

// Clone new Cmd with current config
func (c *Cmd) Clone() *Cmd {
	return NewCommand(c.Bash)
}

// WithTimeout sets command timeout in seconds
func WithTimeout(td int) optionFunc {
	if td < 0 {
		panic("timeout < 0")
	}

	return func(o *Cmd) error {
		o.timeout = td
		return nil
	}
}

// WithShellMode sets shell mode for command
func WithShellMode() optionFunc {
	return func(o *Cmd) error {
		o.ShellMode = true
		return nil
	}
}

// WithExecMode set exec mode, example:
//  ["curl", "-i", "-v", "xiaorui.cc"]
func WithExecMode(b bool) optionFunc {
	return func(o *Cmd) error {
		o.ShellMode = false
		return nil
	}
}

// WithSetDir sets working directory
func WithSetDir(dir string) optionFunc {
	return func(o *Cmd) error {
		o.Dir = dir
		return nil
	}
}

// WithSetEnv sets environment variables
func WithSetEnv(env []string) optionFunc {
	return func(o *Cmd) error {
		o.Env = env
		return nil
	}
}

func (c *Cmd) buildCtx() {
	if c.timeout > 0 {
		c.ctx, c.cancel = context.WithTimeout(context.Background(), time.Duration(c.timeout)*time.Second)
	} else {
		c.ctx, c.cancel = context.WithCancel(context.Background())
	}
}

func (c *Cmd) run() error {
	var (
		cmd *exec.Cmd

		sysProcAttr *syscall.SysProcAttr
	)

	c.buildCtx()

	sysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	c.Status.startTime = time.Now()
	if c.ShellMode {
		cmd = exec.Command(whichsh, "-c", c.Bash)
	} else {
		args := strings.Split(c.Bash, " ")
		cmd = exec.Command(args[0], args[1:]...)
	}

	cmd.Dir = c.Dir
	cmd.Env = c.Env
	cmd.SysProcAttr = sysProcAttr

	// merge multi writer
	mergeStdout := io.MultiWriter(&c.output, &c.stdout)
	mergeStderr := io.MultiWriter(&c.output, &c.stderr)

	// reset writer
	cmd.Stdout = mergeStdout
	cmd.Stderr = mergeStderr
	c.stdcmd = cmd

	// async start
	err := c.stdcmd.Start()
	if err != nil {
		c.Status.Error = err
		return err
	}

	go c.handleWait()

	return nil
}

func (c *Cmd) handleWait() error {
	defer func() {
		if c.Status.Finish {
			return
		}
		c.statusChan <- c.Status
		c.finalize()
	}()

	c.handleTimeout()

	// join process
	err := c.stdcmd.Wait()
	if c.ctx.Err() == context.DeadlineExceeded {
		return err
	}
	if c.ctx.Err() == context.Canceled {
		return err
	}

	if err != nil {
		c.Status.Error = formatExitCode(err)
		return err
	}

	c.Status.Stdout = c.stdout.String()
	c.Status.Stderr = c.stderr.String()
	c.Status.Output = c.output.String()
	return nil
}

// handleTimeout if using commandContext timeout.
// Will not work in shell mode.
func (c *Cmd) handleTimeout() {
	if c.timeout <= 0 {
		return
	}

	call := func() {
		select {
		case <-c.doneChan:
			// safe exit

		case <-c.ctx.Done():
			if c.ctx.Err() == context.Canceled {
				c.Status.Error = ErrProcessCancel
			}
			if c.ctx.Err() == context.DeadlineExceeded {
				c.Status.Error = ErrProcessTimeout
			}
			c.Stop()
		}
	}

	time.AfterFunc(time.Duration(c.timeout)*time.Second, call)
}

func (c *Cmd) finalize() {
	c.Lock()
	defer c.Unlock()

	if c.isFinalized {
		return
	}

	c.Status.CostTime = time.Since(c.Status.startTime)
	c.Status.Finish = true
	c.Status.PID = c.stdcmd.Process.Pid
	c.Status.ExitCode = c.stdcmd.ProcessState.ExitCode()

	// notify
	close(c.doneChan)
	close(c.statusChan)
	c.isFinalized = true
}

// Stop kill -9 pid
func (c *Cmd) Stop() {
	if c.stdcmd == nil || c.stdcmd.Process == nil {
		return
	}

	c.cancel()
	c.finalize()
	c.stdcmd.Process.Kill()
	syscall.Kill(-c.stdcmd.Process.Pid, syscall.SIGKILL)
}

// Kill sends a custom signal to the process
func (c *Cmd) Kill(sig syscall.Signal) {
	syscall.Kill(c.stdcmd.Process.Pid, sig)
}

// Cost returns the time cost for the process
func (c *Cmd) Cost() time.Duration {
	return c.Status.CostTime
}

func formatExitCode(err error) error {
	if err == nil {
		return err
	}

	if strings.Contains(err.Error(), "exit status 127") {
		return ErrNotFoundCommand
	}
	if strings.Contains(err.Error(), "exit status 126") {
		return ErrNotExecutePermission
	}
	if strings.Contains(err.Error(), "exit status 128") {
		return ErrInvalidArgs
	}

	return err
}

// LookPath checks if a command is in the PATH
func LookPath(cmd string) bool {
	if _, err := exec.LookPath(cmd); err != nil {
		return false
	}
	return true
}

// CheckPnameRunning checks to see if a program
// is running.
func CheckPnameRunning(pname string) bool {
	out, _, _ := CommandFormat("ps aux | grep %s |grep -v grep", pname)
	return strings.Contains(out, pname)
}

func logTimes(cmd *exec.Cmd) error {
	ts := int(cmd.ProcessState.SystemTime())
	tu := int(cmd.ProcessState.UserTime())
	cmd.ProcessState.SysUsage()

	// TODO format and log these if option is on ...
	_ = ts
	_ = tu
	return nil
}

// Command executes a command and returns
// CombinedOutput, exitcode, and err.
// If Stderr contains a message, the Go error
// is wrapped in a Stderr message.
func Command(args string) (out string, errno int, err error) {
	defer processErrNo(&errno, err)
	cmd := exec.Command(whichsh, "-c", args)
	b, err := cmd.CombinedOutput()
	logTimes(cmd)

	errno = cmd.ProcessState.ExitCode()
	out = string(b)
	return
}

// Command executes a formatted command and returns
// CombinedOutput, exitcode, and err
func CommandFormat(format string, vals ...interface{}) (string, int, error) {
	s := fmt.Sprintf(format, vals...)
	return Command(s)
}

// CommandContainsAll executes a command then searches
// for matches to any substrings in the output string.
// If ANY of the substrings are not found, the command
// returns false.
func CommandContainsAll(args string, subs ...string) bool {
	b, _, err := Command(args)
	if err != nil {
		return false
	}

	out := string(b)
	for _, sub := range subs {
		if !strings.Contains(out, sub) {
			return false
		}
	}
	return true
}

// CommandScript writes a script to a temp file and
// executes that script.
func CommandScript(script []byte) (string, int, error) {
	fpath := fmt.Sprintf("/tmp/go-shell-%s", randString(16))
	defer os.RemoveAll(fpath)

	err := ioutil.WriteFile(fpath, script, 0666)
	if err != nil {
		return "", DefaultExitCode, errors.Errorf("dump script to file failed, err: %s", err.Error())
	}

	out, code, err := CommandFormat(shFMT, fpath)
	return out, code, err
}

// CommandWithOutErr runs command and return separate
// output: string(stdout), string(stderr), exitcode, err
func CommandWithOutErr(cmd string) (string, string, int, error) {
	var (
		stdout, stderr bytes.Buffer
		err            error
	)

	runner := exec.Command(whichsh, "-c", cmd)
	runner.Stdout = &stdout
	runner.Stderr = &stderr
	err = runner.Start()
	if err != nil {
		return stdout.String(), stderr.String(), runner.ProcessState.ExitCode(), err
	}

	err = runner.Wait()
	return stdout.String(), stderr.String(), runner.ProcessState.ExitCode(), err
}

// CommandWithChan executes the command and returns
// results in a channel
func CommandWithChan(cmd string, queue chan string) error {
	runner := exec.Command("bash", "-c", cmd)
	stdout, err := runner.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err := runner.StderrPipe()
	if err != nil {
		return err
	}

	runner.Start()

	call := func(in io.ReadCloser) {
		reader := bufio.NewReader(in)
		for {
			line, _, err := reader.ReadLine()
			if err != nil || io.EOF == err {
				break
			}

			select {
			case queue <- string(line):
			default:
			}
		}
	}

	go call(stdout)
	go call(stderr)

	runner.Wait()
	close(queue)
	return nil
}

type OutputBuffer struct {
	buf   *bytes.Buffer
	lines []string
	*sync.Mutex
}

func NewOutputBuffer() *OutputBuffer {
	out := &OutputBuffer{
		buf:   &bytes.Buffer{},
		lines: []string{},
		Mutex: &sync.Mutex{},
	}
	return out
}
func (rw *OutputBuffer) Write(p []byte) (n int, err error) {
	rw.Lock()
	n, err = rw.buf.Write(p) // and bytes.Buffer implements io.Writer
	rw.Unlock()
	return
}

func (rw *OutputBuffer) Lines() []string {
	rw.Lock()
	s := bufio.NewScanner(rw.buf)
	for s.Scan() {
		rw.lines = append(rw.lines, s.Text())
	}
	rw.Unlock()
	return rw.lines
}

type OutputStream struct {
	streamChan chan string
	bufSize    int
	buf        []byte
	lastChar   int
}

// NewOutputStream creates a new streaming output on the given channel.
func NewOutputStream(streamChan chan string) *OutputStream {
	out := &OutputStream{
		streamChan: streamChan,
		bufSize:    16384,
		buf:        make([]byte, 16384),
		lastChar:   0,
	}
	return out
}

// Write makes OutputStream implement the io.Writer interface.
func (rw *OutputStream) Write(p []byte) (n int, err error) {
	n = len(p) // end of buffer
	firstChar := 0

	for {
		newlineOffset := bytes.IndexByte(p[firstChar:], '\n')
		if newlineOffset < 0 {
			break // no newline in stream, next line incomplete
		}

		// End of line offset is start (nextLine) + newline offset. Like bufio.Scanner,
		// we allow \r\n but strip the \r too by decrementing the offset for that byte.
		lastChar := firstChar + newlineOffset // "line\n"
		if newlineOffset > 0 && p[newlineOffset-1] == '\r' {
			lastChar -= 1 // "line\r\n"
		}

		// Send the line, prepend line buffer if set
		var line string
		if rw.lastChar > 0 {
			line = string(rw.buf[0:rw.lastChar])
			rw.lastChar = 0 // reset buffer
		}
		line += string(p[firstChar:lastChar])
		rw.streamChan <- line // blocks if chan full

		// Next line offset is the first byte (+1) after the newline (i)
		firstChar += newlineOffset + 1
	}

	if firstChar < n {
		remain := len(p[firstChar:])
		bufFree := len(rw.buf[rw.lastChar:])
		if remain > bufFree {
			var line string
			if rw.lastChar > 0 {
				line = string(rw.buf[0:rw.lastChar])
			}
			line += string(p[firstChar:])
			err = ErrLineBufferOverflow
			n = firstChar
			return // implicit
		}
		copy(rw.buf[rw.lastChar:], p[firstChar:])
		rw.lastChar += remain
	}

	return // implicit
}

func (rw *OutputStream) Lines() <-chan string {
	return rw.streamChan
}

func (rw *OutputStream) SetLineBufferSize(n int) {
	rw.bufSize = n
	rw.buf = make([]byte, rw.bufSize)
}
