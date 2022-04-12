package types

import (
	"fmt"
	"os"

	"github.com/mattn/go-isatty"
)

// IsTerminal returns true if os.Stdout is a terminal.
// This is used to determine output options such as
// ANSI color sequences used in terminal output.
//
// Reference: uses isatty package (MIT License):
// https://github.com/mattn/go-isatty
var IsTerminal bool = isTerminal()

func isTerminal() bool {
	return isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
}

// TerminalExample prints the results of isatty.IsTerminal()
// and isatty.IsCygwinTerminal() to os.Stdout
func TerminalExample() {
	fmt.Println("Terminal test using")
	fmt.Println(" Reference: uses isatty package (MIT License):")
	fmt.Println(" https://github.com/mattn/go-isatty")
	fmt.Println("")

	if isatty.IsTerminal(os.Stdout.Fd()) {
		fmt.Println("Is Terminal")
	} else if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		fmt.Println("Is Cygwin/MSYS2 Terminal")
	} else {
		fmt.Println("Is Not Terminal")
	}
}
