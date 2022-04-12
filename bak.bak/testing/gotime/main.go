package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/skeptycal/goutil/repo/errorlogger"
)

var log = errorlogger.New()

func Shell(s ...string) string {
	com := strings.TrimSpace(strings.Join(s, " "))
	command := strings.Split(com, " ")
	if len(command) < 1 {
		return ""
	}
	app := command[0]
	var args []string
	if len(command) < 2 {
		args = []string{""}
	} else {
		args = command[1:]
	}

	cmd := exec.Command(app, args...)
	b, err := cmd.Output()
	if err != nil {
		log.Error(err)
		return err.Error()
	}
	return string(b)
}

func main() {
	// stuff

	prompt := "c:> "

	reader := bufio.NewReader(os.Stdin)

	for {
		// n, err := fmt.Scanln(args)
		fmt.Print(prompt)
		text, _ := reader.ReadString('\n')

		fmt.Println(Shell(text))
	}
}
