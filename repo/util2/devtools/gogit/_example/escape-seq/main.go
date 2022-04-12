package main

import (
	"fmt"

	ansi "
)

func main() {
	// stdOut := bufio.NewWriter(colorable.NewColorableStdout())
	s := ansi.Output
	stdOut := ansi.Output
	ansi.NewStdout(w io.Writer)
	defer stdOut.Flush() // Write buffered writes before exiting

	fmt.Fprint(stdOut, "\x1B[3GMove to 3rd Column\n")
	fmt.Fprint(stdOut, "\x1B[1;2HMove to 2nd Column on 1st Line\n")
}
