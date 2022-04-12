package main

import (
	"os"

	"
)

var r = redlogger.New(os.Stderr)

func main() {
    defer r.Flush()
    r.WriteString("Hello World!")
}
