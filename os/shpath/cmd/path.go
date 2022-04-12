package main

import (
	"fmt"

	"github.com/skeptycal/goutil/os/shpath"
)

func main() {
	sh := shpath.NewPath()
	_ = sh.Clean()

	fmt.Println(sh)
	fmt.Println(sh.Out())

}
