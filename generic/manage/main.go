package main

import (
	"errors"
	"fmt"
	"strings"

	"
	"
)

var (
	log   = errorlogger.Log
	Err   = log.Err
	sh    = ghshell.Command
	GoBot = ghshell.GoBot
)

func main() {

	// gnuflags.GetOptLong()

	// ghshell.Example()
	// gnuflags.Example()

	fmt.Println(ghshell.GitCurrentBranch())

	// gitit("")

	// load config
	// check flags

	// loop
	// 		menu
	// 		call stuff
	// 		check for errors
	// end loop
}

func echo(args ...string) (string, int, error) {
	arglist := strings.Join(args, " ")
	out, errno, err := sh(arglist)
	if err != nil {
		s := fmt.Sprintf("Shell command error (%v) running %v", err, arglist)
		log.Infof("out: %s", out)
		log.Infof("errno: %v", errno)
		log.Infof("err: %v", err)
		if errno != 0 {
			s += fmt.Sprintf(" (%v)", errno)
		}
		Err(errors.New(s))
		return out, errno, err
	}

	if GoBot.Verbose {
		fmt.Println(out)
	}
	return out, errno, err
}

func zip(args ...string) string {
	arglist := strings.Join(args, " ")
	out, errno, err := sh(arglist)
	if err != nil {
		s := fmt.Sprintf("Shell command error (%v) running %v", err, arglist)
		if errno != 0 {
			s += fmt.Sprintf(" (%v)", errno)
		}
		Err(errors.New(s))
		return s
	}
	return out
}
