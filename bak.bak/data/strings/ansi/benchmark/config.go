package benchmark

import (
	"bufio"
	"io"
	"os"
)

// fmtTestInfo is a format string for printing benchmark descriptions
// strconv may be faster ... but likely not needed ...
// Reference: https://stackoverflow.com/questions/38552803/how-to-convert-a-bool-to-a-string-in-go
const fmtTestInfo = "%s:15 %v:15.15 %v:15.15 %t"

const defaultFileMode = 0644

type logFile struct {
	// name of log file; use "" to assign random name
	name string

	// returns the io.Writer for the log file
	file *bufio.Writer

	// cache for file information
	fi os.FileInfo

	// returns the file handle for the log file
	f *os.File

	// set to append to the log (default overwrite)
	append bool

	// maxlogsize sets the maximum file size for the log file
	maxlogsize int64
}

func (f *logFile) Stat() (fi os.FileInfo, err error) {
	if f.fi == nil {
		f.fi, err = os.Stat(f.name)
		if err != nil {
			f.fi = nil
		}
	}
	return f.fi, nil
}

func (f *logFile) File() (ff *os.File, err error) {

	if f.file == nil {
		_, err := f.Stat()
		if err != nil {
			if os.IsNotExist(err) {
				ff, err = os.OpenFile(
					f.name,
					os.O_CREATE|os.O_APPEND,
					defaultFileMode,
				)
				if err != nil {
					return nil, err
				}

				f.fi, err = os.Stat(f.name)
				if err != nil {
					return nil, err
				}

				return f.f, nil
			}
			return nil, err
		}
	}
	return f.f, nil
}

func (f *logFile) Writer() *bufio.Writer {
	if f.file == nil {
		w, err := f.File()
		if err != nil {
			return nil
		}
		f.file = bufio.NewWriter(w)
	}
	return f.file
}

func NewLogFile(name string, append bool, maxlogsize int64) *logFile {
	return &logFile{}
}

type ConfigSettings struct {

	// verbose sets the level of detail to high
	// (for debugging)
	verbose bool

	// output is the io.Writer used to write
	// visual output
	//
	// set to <nil> to use default (os.Stderr)
	output *bufio.Writer

	// saveFile is used to describe logging to disk
	//
	// set to <nil> to disable logging
	log *logFile

	// format string used to print test info
	fmtTestInfo string
}

var defaultConfig *ConfigSettings = NewConfig(true, nil, nil, "")

func NewConfig(verbose bool, out io.Writer, log *logFile, fmtString string) *ConfigSettings {

	if fmtString == "" {
		fmtString = fmtTestInfo
	}

	return &ConfigSettings{
		verbose:     verbose,
		output:      NewWriter(out),
		log:         log,
		fmtTestInfo: fmtString,
	}
}
