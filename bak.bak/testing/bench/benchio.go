package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func AppArgs(in ...string) (app string, args []string) {
	switch len(in) {
	case 0:
		return "", []string{""}
	case 1:
		return in[0], []string{""}
	case 2:
		return in[0], []string{in[1]}
	default:
		return in[0], in[1:]
	}
}

func Shell(in ...string) (string, error) {
	s := strings.TrimSpace(strings.Join(in, " "))
	in = strings.Split(s, " ")
	app, args := AppArgs(in...)

	cmd := exec.Command(app, args...)
	b, err := cmd.Output()

	if err != nil {
		return err.Error(), err
	}

	return string(b), nil
}

func grep(list []string, needle string) []string {
	retval := make([]string, len(list))

	for _, s := range list {
		if strings.Contains(s, needle) {
			retval = append(retval, s)
		}
	}

	return retval
}

func main() {
	buildio, err := Shell("go build -gcflags='m -m' io")
	if err != nil {
		log.Fatal(err)
	}

	list := strings.Split(buildio, "\n")

	fmt.Println(grep(list, "escapes to heap"))

}
