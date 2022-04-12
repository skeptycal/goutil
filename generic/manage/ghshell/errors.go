package ghshell

import (
	"fmt"

	"github.com/pkg/errors"
)

type shellError struct {
	errno   int
	err     error
	msg     string
	command string
}

func (e *shellError) ErrNo() int { return e.errno }

func (e *shellError) Error() string {
	if e.msg == "" {
		e.msg = "shell command error"
	}
	if e.command == "" {
		e.command = "shell command"
	}
	return fmt.Sprintf("%s: %q(%v) during %q", e.msg, e.errno, e.err, e.command)
}

func (e *shellError) String() string {
	return e.Error()
}

func (e *shellError) Is(target error) bool {
	return errors.Is(e.err, target)
}

func (e *shellError) As(target any) bool {
	return errors.As(e.err, target)
}

func (e *shellError) Wrap(msg string) error {
	e.err = errors.Wrap(e.err, msg)
	return e.err
}

func NewShellError(errno *int, err error, msg *string, command *string) *shellError {
	if processErrNo(errno, err) == nil {
		return nil
	}

	if *msg == "" {
		*msg = "shell command error"
	}
	if *command == "" {
		*command = "shell command"
	}

	return &shellError{
		errno:   *errno,
		err:     err,
		msg:     *msg,
		command: *command,
	}
}

func processErrNo(errno *int, err error) error {
	if *errno == 0 {
		if err == nil {
			return nil
		}
		return errors.Wrap(err, fmt.Sprintf("shell command error without errno: %v", err))
	}
	if err == nil {
		return errors.New(fmt.Sprintf("shell command error without Go error: %v", errno))
		// fmt.Errorf("shell command error without Go error: %v", errno)
	}
	return errors.Wrap(err, fmt.Sprintf("shell command error: %v(%v)", errno, err))
}

var (
	ErrLineBufferOverflow = errors.New("line buffer overflow")

	ErrAlreadyFinished      = errors.New("already finished")
	ErrNotFoundCommand      = errors.New("command not found")
	ErrNotExecutePermission = errors.New("not execute permission")
	ErrInvalidArgs          = errors.New("Invalid argument to exit")
	ErrProcessTimeout       = errors.New("throw process timeout")
	ErrProcessCancel        = errors.New("active cancel process")

	DefaultExitCode = 2
)
