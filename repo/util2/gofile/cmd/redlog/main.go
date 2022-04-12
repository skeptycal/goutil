package main

import (
	"os"

	"github.com/skeptycal/goutil/os/redlogger"
)

var r = redlogger.New(os.Stderr)

func main() {
	defer r.Flush()
	r.WriteString("Hello World!")
}
