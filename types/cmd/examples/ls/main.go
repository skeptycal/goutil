package main

import (
	"flag"
	"time"

	"github.com/sirupsen/logrus"
	"
	"
)

const (
	defaultFilesArg = "."
)

var (
	log     = errorlogger.Log
	blue    = types.Blue
	attn    = types.Attn
	verbose = types.Info

	verboseFlag  bool
	logLevelFlag int
	filesArg     string
)

func init() {
	flag.BoolVar(&verboseFlag, "v", false, "verbose, detailed log output")
	flag.BoolVar(&verboseFlag, "verbose", false, "verbose, detailed log output")
	flag.IntVar(&logLevelFlag, "l", 1, "log output level of detail")
	flag.IntVar(&logLevelFlag, "loglevel", 1, "log output level of detail")
	flag.StringVar(&filesArg, "f", defaultFilesArg, "file pattern to search for")
	flag.StringVar(&filesArg, "files", defaultFilesArg, "file pattern to search for")

	flag.Parse()
}

const (
	fmtString = " %s %8.0d %20v %-20v \n"
	fmtTime   = time.ANSIC
)

func main() {
	if verboseFlag {
		log.SetLevel(logrus.InfoLevel)
		log.Info("Verbose output enabled...")
	} else {
		log.SetLevel(logrus.ErrorLevel)
	}
	suffix := ".go"

	list, err := types.GetFileListBySuffix(filesArg, suffix)
	if err != nil {
		log.Fatalf("error reading directory listing: %v\n", err)
	}

	verbose("filesArg: %s", filesArg)
	verbose("pattern: %s", suffix)
	verbose("list: %v", list)
	verbose("number of files: %v", len(list))

	for _, name := range list {
		verbose(name)
		attn(name)
		blue(name)
	}
}

// s := fmt.Sprintf(fmtString, fi.Mode(), fi.Size(), fi.ModTime().Format(fmtTime), name)
