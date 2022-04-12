package find

import (
	"filepath"
	"os/exec"
	"strings"
)

// Stat is exported out of golang convention, rather than necessity
// type Stat struct {
// 	Name string
// 	Path string
// }

// Repo uses git via the console to locate the top level directory
// func Repo() (Stat, error) {
// 	path, err := rootPath()
// 	if err != nil {
// 		return Stat{
// 			"Unknown",
// 			"./",
// 		}, err
// 	}

// 	gitRepo, err := exec.Command("basename", path).Output()
// 	if err != nil {
// 		return Stat{}, err
// 	}

// 	return Stat{
// 		strings.TrimSpace(string(gitRepo)),
// 		path,
// 	}, nil
// }

func AppArgs(s string) (app string, args string) {
	list := strings.Split(s, " ")
	switch len(list) {
	case 0:
		return "", ""
	case 1:
		return list[0], ""
	case 2:
		return list[0], list[1]
	default:
		return list[0], strings.Join(list[1:], " ")
	}
}

func Shell(command string) (string, error) {
	app, args := AppArgs(command)
	path, err := filepath.Abs(app)
	if err != nil {
		return "", err
	}
	cmd := exec.Command(app, args)
	cmd.Run()
	b, err := cmd.Output()
	out := strings.TrimSpace(string(b))
	if err != nil {

		return "", err
	}
	return out, nil
}

// func rootPath() (string, error) {
// 	path, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
// 	if err != nil {
// 		return "", err
// 	}
// 	return strings.TrimSpace(string(path)), nil
// }
