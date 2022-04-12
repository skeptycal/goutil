package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: reverse [file ...]")
		os.Exit(0)
	}

	// arg := os.Args[1]

}
