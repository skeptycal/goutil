package main

// ansistrings demo

import (
	"fmt"

	"github.com/skeptycal/goutil/data/strings/ansi"
)

func main() {

	s := "This is a string."

	fmt.Println(s)

	a := ansi.NewColor("172", "0", "1")

	fmt.Print(a)
	fmt.Println(s)
}
