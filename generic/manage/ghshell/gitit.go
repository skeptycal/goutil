package ghshell

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/pkg/errors"
)

var (
	sh          = Command
	noFilesList = []fs.DirEntry{}
)

const defaultMessage = `Gobot: dev progress update`

func Message(msg string) string {
	if msg == "" {
		return defaultMessage
	}
	return msg
}

func GitCommit(message string) error {
	if message == "" {
		message = GoBot.GobotDevAutosaveMessage
	}

	// TODO message function is not implemented
	return shErr("git commit -m '" + Message(message) + "'")
}

func GitCurrentBranch() (out string, err error) { return GetOutput("git branch --show-current") }
func GitAddAll() error                          { return shErr("git add --all") }
func GitPush() error                            { return shErr("git push") }
func GitPushSetUpstreamOrigin() error {
	gcb, err := GitCurrentBranch()
	if err != nil {
		return err
	}
	return shErr("git push --set-upstream origin " + gcb)
}
func GitFetch() error { return shErr("git fetch") }

// GitCommitAll processes the following commands
// in sequence:
//  git add -all
//  git commit -m "message" (or default)
//  git push origin main
func GitCommitAll(message string) error {
	err := GitAddAll()
	if err != nil {
		return err
	}

	GitCommit(message)

	return shErr("git push")
}

func (gb goBot) GetGoFiles() ([]fs.DirEntry, error) {

	list, err := os.ReadDir(gb.PWD())
	if err != nil {
		return noFilesList, err
	}

	// assume half of the files are go files ...
	out := make([]fs.DirEntry, 0, len(list)/2)

	for _, file := range list {
		if !strings.HasSuffix(file.Name(), ".go") {
			out = append(out, file)
		}
	}
	return out, nil
}

// Gitit is a quick and dirty git repo
// autosave and push.
//
// Gitit runs the following commands in order:
//  go mod tidy
//  go doc >|go.doc
//  git add --all
//  git commit -m $message
//  git push
func Gitit(message string) error {
	// path := gofile.PWD()
	// var goDir bool = true

	// goDir := gb.IsGoDir(path)

	var goDir bool

	if goDir {
		err := shErr("go mod tidy")
		if err != nil {
			return fmt.Errorf("unable to run 'go mod tidy': %v", err)
		}
		err = shErr("go doc >|go.doc")
		if err != nil {
			return fmt.Errorf("unable to run 'go doc >|go.doc': %v", err)
		}
	}
	GitCommitAll(message)

	// replaces shell script:
	// 	go mod tidy && go doc >|go.doc
	// git add --all
	// git commit -m "${1:-'GitBot: dev progress autosave'}"
	// git push --set-upstream origin $(git_current_branch)

	return nil
}

// shErr processes the shell command and
// returns only the error. If no error was
// encountered, nil is returned.
//
// It uses Echo to process the command, so
// any output or logging directives are handled.
func shErr(args ...string) error {
	_, err := Echo(args...)
	return err
}

// Echo processes the shell command and
// and returns the result and any error
// encountered. In addition, it prints the
// result to stdout and any error message
// to stderr.
//
// Any shell error values (errno of type int)
// are wrapped into the Go error message.
//
// The io.Writers for stdout and stderr
// can be redirected in the configuration
// and all stdout and stderr output can be
// enabled or disabled by using the config
// settings Verbose and Debug, respectively.
func Echo(args ...string) (out string, err error) {
	errno := 0
	arglist := strings.Join(args, " ")
	out, errno, err = sh(arglist)
	err = processErrNo(&errno, err)

	GoBot.Println(out)

	return out, err
}

// GetOutput returns the contents of the Stdout.
// Any error is returned wrapped with a Stderr
// message and any errno returned from the shell.
func GetOutput(args ...string) (out string, err error) {
	var errno = 0
	arglist := strings.Join(args, " ")
	out, errno, err = sh(arglist)
	processErrNo(&errno, err)
	errors.Wrap(err, fmt.Sprintf("Shell command error: (%v) - (%v) running %v", errno, err, arglist))
	return
}
