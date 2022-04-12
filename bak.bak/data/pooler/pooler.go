package pooler

import (
	"log"
	"os"
)

const envDebugName = "GO_DEBUG"

var debug bool = false

func init() {
	if os.Getenv(envDebugName) == "true" {
		debug = true
	}
	dbLog("debug printing is active ...")
}

// dbLog prints unformatted log messages if the variable
// 'debug' is true. The default is false. 'debug' can be set
// by creating an environment variable named "GO_DEBUG"
// and setting the value to "true".
func dbLog(args ...interface{}) {
	if debug {
		log.Print(args...)
	}
}

// dbLogf prints formatted log messages if the variable
// 'debug' is true. The default is false. 'debug' can be set
// by creating an environment variable named "GO_DEBUG"
// and setting the value to "true".
func dbLogf(format string, args ...interface{}) {
	if debug {
		log.Printf(format, args...)
	}
}
