package ghshell

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"

	"
	"
)

const (
	defaultAutoSaveInterval   = 10 * time.Minute
	defaultDevAutosaveMessage = `GoBot: dev progress autosave`
)

var (
	GoBot = goBot{
		//* State variables
		currentMessage: "",
		currentDir:     gofile.PWD(),

		//* Defaults
		DefaultGitCommitMessage: "",
		GobotDevAutosaveMessage: defaultDevAutosaveMessage,
		GobotAutoSaveTimer:      defaultAutoSaveInterval,
	}

	GoBotCliWriter = cliWriter{
		Writer:  nil,
		Verbose: true,
		Out:     os.Stdout,
		Debug:   true,
		Err:     os.Stderr,
	}
)

type (
	cliWriter struct {
		//////////////// Default I/O:

		Writer io.Writer

		// Verbose indicates whether details of
		// commands are echoed to stdout.
		Verbose bool

		// Out is the Writer assigned to normal
		// output messages. (StdOut)
		Out io.Writer

		// Debug indicates whether debug messages
		// containing errors or other debug information
		// are echoed to stdout.
		Debug bool

		// Err is the Writer assigned to dev, debug,
		// or error output messages. (StdErr)
		Err io.Writer

		// Islogging indicates whether logging is enabled.
		IsLogging bool

		// Log is the configured logger.
		Log errorlogger.ErrorLogger
	}

	goBot struct {
		currentMessage string
		currentDir     string
		repoRoot       string
		env            map[string]string

		cliWriter

		//////////////// Default Strings:

		// GobotDevAutosaveMessage is the message
		// used for automatic commits when no other
		// message is supplied.
		GobotDevAutosaveMessage string

		DefaultGitCommitMessage string

		// GobotAutoSaveTimer is the interval
		// between automatic dev progress commits.
		// A value of 0 means automatic progress
		// uses the default value of 10 min.
		GobotAutoSaveTimer time.Duration
	}
)

// TODO idea from exec_posix
//
// addCriticalEnv adds any critical environment variables that are required
// (or at least almost always required) on the operating system.
// Currently this is only used for Windows.
func addCriticalEnv(env []string) []string {
	if runtime.GOOS != "windows" {
		return env
	}
	for _, kv := range env {
		k, _, ok := strings.Cut(kv, "=")
		if !ok {
			continue
		}
		if strings.EqualFold(k, "SYSTEMROOT") {
			// We already have it.
			return env
		}
	}
	return append(env, "SYSTEMROOT="+os.Getenv("SYSTEMROOT"))
}

func (gb goBot) loadEnv() {
	// TODO load the Environment
	// check for critical environment variables
	// add default critical env if available
	// add default values (for GOOS) if unable to determine them
}

// GoDir indicates whether the current
// working directory contains Go files.
//
// If there is an error reading the
// working directory or getting a list
// of files, an error is logged.
func (gb goBot) IsGoDir() bool {
	var err error
	defer de(&err)

	files, err := GoBot.GetGoFiles()
	if err != nil {
		err = fmt.Errorf("(gb.IsGoDir) error getting list of go files: %v", err)
	}

	if len(files) < 1 {
		err = fmt.Errorf("no go files found in %v", gb.PWD())
		return false
	}

	return false
}

func (gb goBot) GitCurrentBranch() (out string, err error) {
	if gb.repoRoot == "" {
		gb.repoRoot, err = GetOutput("git branch --show-current")
		if err != nil {
			gb.repoRoot = ""
		}
	}
	return gb.repoRoot, err
}

func (gb goBot) GoDir() string {
	if gb.currentDir == "" {
		gb.currentDir = gb.PWD()
	}
	return gb.currentDir
}

func (gb goBot) Message() string {
	_ = gb.DefaultGitCommitMessage
	if gb.currentMessage == "" {
		gb.currentMessage = gb.defaultMessage()
	}
	return gb.currentMessage
}

func (gb goBot) defaultMessage() string {
	if gb.DefaultGitCommitMessage == "" {
		gb.currentMessage = defaultDevAutosaveMessage
		_ = gb.currentMessage // TODO go vet message??
	}
	return gb.DefaultGitCommitMessage
}

func (gb goBot) FindGoModPath() string {
	// TODO ... implement this ...
	return ""
}

func (gb goBot) PWD() string {
	if gb.currentDir == "" {
		gb.currentDir = gofile.PWD()
	}
	return gb.currentDir
}

func (gb goBot) setCommitMessage(message string) {
	if message == "" {
		gb.DefaultGitCommitMessage = gb.GobotDevAutosaveMessage
		_ = gb.DefaultGitCommitMessage // TODO go vet message??
	}
	gb.DefaultGitCommitMessage = message
	_ = gb.DefaultGitCommitMessage // TODO go vet message??

}

func (gb goBot) setWriters() io.Writer {
	writerlist := make([]io.Writer, 0, 3)
	if gb.Verbose {
		writerlist = append(writerlist, gb.Out)
	}
	if gb.Debug {
		writerlist = append(writerlist, gb.Err)
	}
	if gb.IsLogging {
		writerlist = append(writerlist, gb.errorloggerWriter())
	}
	gb.Writer = io.MultiWriter(writerlist...)
	return gb.Writer
}

func (gb goBot) errorloggerWriter() io.Writer {
	// TODO new errorlogger method Writer() io.Writer was added ... update dependencies and test
	// return gb.Log.Writer
	return nil
}

func (gb goBot) Println(a ...interface{}) {
	if gb.Verbose {
		fmt.Fprintln(gb.Out, a...)
	}
}

func (gb goBot) Printf(format string, a ...interface{}) {
	if gb.Writer == nil {
		gb.setWriters()
	}

	fmt.Fprintf(gb.Writer, format, a...)

	if gb.Debug {
		fmt.Fprintf(gb.Out, format, a...)
	}
	if gb.IsLogging {
		gb.Log.Printf(format, a...)
	}
}
