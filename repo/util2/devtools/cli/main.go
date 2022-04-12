package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"
)

func main() {
	var (
		cmd          *exec.Cmd
		background   context.Context = context.Background()
		app          string
		args         []string
		promptString string = gofile.PWD() + "\n➜ "
	)

	for {
		rin := bufio.NewReader(os.Stdin)
		fmt.Print(promptString)
		text, _ := rin.ReadString('\n')
		fmt.Println(text)

		arglist := strings.Fields(text)

		app = arglist[0]
		if len(arglist) > 1 {
			args = arglist[1:]
		} else {
			args = []string{""}
		}

		cmd = exec.CommandContext(background, app, args...)

		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		err = cmd.Wait()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(cmd.Stdout)
	}
}
