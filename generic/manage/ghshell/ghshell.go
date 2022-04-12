package ghshell

import (
	"fmt"

	"
)

var (
	log = errorlogger.New()
	Err = log.Err
)

func init() {
	log.SetLevel(errorlogger.InfoLevel)
}

var (
	TES             = ""
	NL              = "\n"
	TAB             = "\t"
	emptyStringList = []string{""}
)

func Example() {
	fmt.Println("Example import verification from goutil/manage/ghshell")
}

// func out(app string, args ...string) string {
// 	b, err := exec.Command(app, args...).Output()
// 	if err != nil {
// 		return err.Error()
// 	}
// 	return string(b)
// }

// func combined(app string, args ...string) (string, error) {
// 	b, err := exec.Command(app, args...).CombinedOutput()
// 	if err != nil {
// 		return err.Error(), err
// 	}
// 	return string(b), nil
// }

// func Out(args ...string) string {
// 	app, args, err := AppArgs(TES, args...)
// 	if err != nil {
// 		return app
// 	}
// 	return out(app, args...)
// }

// func Combined(args ...string) (string, error) {
// 	app, args, err := AppArgs(TES, args...)
// 	if err != nil {
// 		return app, err
// 	}
// 	return combined(app, args...)
// }

// func Shell(args ...string) (string, error) {
// 	usage := "Usage: Shell(app, args...) (string, error)"
// 	app, args, err := AppArgs(usage, args...)
// 	if err != nil {
// 		return app, err
// 	}

// 	return combined(app, args...)
// }

type commandString string

const (
	cs_gitInit   commandString = `git init`
	cs_ls        commandString = `ls -lah`
	cs_gitstatus commandString = `git status`
	cs_gaa       commandString = `git add --all`
	cs_gitignore commandString = `curl -fLw '\n' https://www.gitignore.io/api/"${(j:,:)@}"`
)

// // func clean(args ...string) []string {

// // 	for _, arg := range args {
// // 		for i, c := range arg {
// // 			if unicode.IsControl(r) || unicode.IsMark(r) {

// // 			}
// // 		}
// // 	}
// // 	args = toFields(args...)
// // 	s := strings.Join(args, " ")
// // 	args = strings.Fields(s)

// // 	// check for invalid characters  ...
// // 	return args
// // }

// // gi sends a request to the gitignore.io api
// // and returns text for a .gitignore file.
// func MakeGitIgnore(envs ...string) string {
// 	s := strings.Join(envs, " ")
// 	args := strings.Fields(s)
// 	s = `curl -fLw '\n' https://www.gitignore.io/api/ `
// 	s += strings.Join(envs, " ")
// 	app, args, _ := AppArgs("", args...)

// 	return out(app, args...)

// 	// args := `curl -fLw '\n' https://www.gitignore.io/api/"${(j:,:)@}"`
// 	// return zip(gitignore)
// }

// // zip is a quick zsh return with a bit less
// // error checking. It is designed for internal
// // use when pretested input strings are used.
// func zip(s commandString) string {
// 	app, args, err := AppArgs("", string(s))
// 	if err != nil {
// 		if app == "" {
// 			return err.Error()
// 		}
// 		return app
// 	}
// 	b, err := exec.Command(app, args...).Output()
// 	return string(b)
// }

// // Zsh runs shell commands in '-c' mode using zsh.
// // arguments provided and returns the combined
// // standard output and standard error.
// //
// // Any error is also returned as a standard error.
// //
// // As a courtesy, a usage string will be returned
// // instead of the default empty
// // string when no arguments are provided.
// func Zsh(args ...string) (string, error) {
// 	list := make([]string, len(args)+2)
// 	list[0] = "zsh"
// 	list[1] = "-c"
// 	for i, arg := range args {
// 		list[i+2] = arg
// 	}

// 	zsh_usage := zip("zsh --help")
// 	zshUsageSlice := strings.Split(zsh_usage, NL)
// 	log.Infof("Usage: %s", zshUsageSlice[0])

// 	zsh_version := zip("zsh --version")

// 	usage := "Usage: Zsh(args...) (string, error)"
// 	app, args, err := AppArgs(usage, list...)

// 	app, args, err := AppArgs(usage, args...)
// 	if err != nil {
// 		return app, err
// 	}

// 	cmd := exec.Command(app, args...)
// 	b, err := cmd.CombinedOutput()
// 	return string(b), err
// }

// AppArgs parses the input arguments into a set of
// app, args strings variables designed for use
// with exec.Command.
//
// Arguments are split on any unicode whitespace and
// the first arg is returned as the 'app' while the
// remaining are returned as a slice of strings.
//
// Any error is also returned as a standard error.
//
// As a courtesy, if a usage string is provided,
// it will be used instead of the default empty
// string 'app' when no arguments are provided.
func AppArgs(usage string, in ...string) (appMessage string, args []string, err error) {
	args = ToFields(in...)
	appMessage = ""
	switch len(args) {
	case 0:
		err = fmt.Errorf("no command specified")
		if usage == "" {
			usage = err.Error()
		}
		return usage, emptyStringList, err
	case 1:
		return args[0], emptyStringList, nil
	default:
		return args[0], args[1:], nil
	}
}
