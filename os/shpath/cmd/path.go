package main

import (
	"fmt"

	"
)

func main() {
	sh := shpath.NewPath()
	_ = sh.Clean()

	fmt.Println(sh)
	fmt.Println(sh.Out())

}
