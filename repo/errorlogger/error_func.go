// Copyright (c) 2021 Michael Treanor
// https://
// MIT License

package errorlogger

import "github.com/pkg/errors"

// Disable disables logging and sets a no-op function for
// Err() to prevent slowdowns while logging is disabled.
func (e *errorLogger) Disable() {
	e.errFunc = e.noErr
}

// Enable enables logging and restores the Err() logging functionality.
func (e *errorLogger) Enable() {
	e.errFunc = e.yesErr
}

// Err logs an error to the provided logger, if it is enabled,
// and returns the error unchanged to be propagated up.
func (e *errorLogger) Err(err error) error {
	if err == nil {
		return nil
	}
	return e.errFunc(err)
}

// noErr is a no-op errorFunc for disabling logging without
// constant repetitive flag checks or other hacks.
// https://en.wikipedia.org/wiki/NOP_(code)
func (e *errorLogger) noErr(err error) error {
	// TODO does this really need to return an error?
	// TODO does the compiler remove this?
	// TODO would a pointer be better here? *(&err)
	// TODO does the compiler remove this?
	return err
}

// yesErr is an errorFunc that logs and wraps an error, then
// returns the error unchanged.
func (e *errorLogger) yesErr(err error) error {
	if err == nil {
		return nil
	}
	if e.wrap != nil {
		err = errors.Wrap(err, e.wrap.Error())
	}
	e.logFunc(err)

	return err
}
