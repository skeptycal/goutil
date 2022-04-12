package main

import (
	"fmt"
	"os"
	"path/filepath"

	"
	"
)

var log = errorlogger.Log

func main() {
	fmt.Printf("%25.25s %s\n", "raw os.Args[0]:", os.Args[0])

	arg0, err := osargsutils.Arg0()
	if err != nil {
		log.Fatal(err)
	}

	arg0 = filepath.Base(arg0)
	fmt.Printf("%25.25s %s\n", "using Arg0():", arg0)

	here, me, err := osargsutils.HereMe()
	if err != nil {
		log.Fatal(err)
	}

	here = filepath.Base(here)

	fmt.Printf("%25.25s - Here: %s Me: %s\n", "using HereMe():", here, me)

}
