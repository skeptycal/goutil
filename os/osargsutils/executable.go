package osargsutils

import (
	"os"
	"path"
	"path/filepath"

	"github.com/skeptycal/goutil/repo/errorlogger"
)

var log = errorlogger.Log
var Err = log.Err

// Arg0 returns the absolute path name for the executable that started the
// current process. EvalSymlinks is run on the resulting path to provide a
// stable result.
func Arg0() (string, error) {
	// As of Go 1.8 (Released February 2017) the recommended
	// way of doing this is with os.Executable.
	return zeroOsExecutable()
}

func HereMe() (string, string, error) {
	// hereMe returns the folder (here) and basename (me) of
	// the executable that started the current process.
	return hereMe()
}

// hereMe returns the folder (here) and basename (me) of
// the executable that started the current process.
func hereMe() (string, string, error) {
	// As of Go 1.8 (Released February 2017) the recommended
	// way of doing this is with os.Executable:
	zero, err := Arg0()
	if err != nil {
		return "", "", Err(err)
	}

	// TODO - using path.Split() returns dir ending
	// with a slash, where Dir() would not
	return filepath.Dir(zero), filepath.Base(zero), nil
}

// hereMe2 returns the folder (here) and basename (me) of
// the executable that started the current process.
func hereMe2() (string, string, error) {
	// As of Go 1.8 (Released February 2017) the recommended
	// way of doing this is with os.Executable:
	zero, err := Arg0()
	if err != nil {
		return "", "", Err(err)
	}

	// TODO - using path.Split() returns dir ending
	// with a slash, where Dir() would not
	dir, base := path.Split(zero)
	return dir, base, nil
}

func zeroOsArgs() (string, error) {
	// Prior to Go 1.8, you could use os.Args[0]
	ex, err := filepath.Abs(os.Args[0])
	if err != nil {
		return "", Err(err)
	}

	return filepath.EvalSymlinks(ex)
}

func zeroOsExecutable() (string, error) {
	// As of Go 1.8 (Released February 2017) the recommended
	// way of doing this is with os.Executable:
	ex, err := os.Executable()
	if err != nil {
		return "", Err(err)
	}

	return filepath.EvalSymlinks(ex)
}

func rawOsArgsZero() (string, error) {
	ex := os.Args[0]
	return filepath.EvalSymlinks(ex)
}
