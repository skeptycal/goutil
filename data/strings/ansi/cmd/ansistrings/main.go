package main

// ansistrings demo

import (
	"fmt"

	"
)

func main() {

	s := "This is a string."

	fmt.Println(s)

	a := ansi.NewColor("172", "0", "1")

	fmt.Print(a)
	fmt.Println(s)
}
