package writers

import (
	"io"

	locker "
)

func NewLockEnableWriter(w ioWriter) LockEnableWriter {
	return &lockEnableWriter{locker.NewLocker(), defaultNopWriter(nil)}
}

type (
	LockEnableWriter interface {
		locker.LockEnabler
		ioWriter
	}

	lockEnableWriter struct {
		locker.LockEnabler
		io.Writer
	}

	// ioWriter implements io.ioWriter
	ioWriter interface {
		Write(p []byte) (n int, err error)
	}

	// ioStringWriter implements io.ioStringWriter
	ioStringWriter interface {
		WriteString(string) (n int, err error)
	}
)
