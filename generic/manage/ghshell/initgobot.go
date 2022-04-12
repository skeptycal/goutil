package ghshell

import "os"

// var criticalEnv = []string{}

func ErrControl() error {
	return nil
}

func Home() string {
	var err error
	defer ErrControl()
	s, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return s
}

// criticalEnv contains a list of environment variables
// and default values that are required to execute shell
// commands with GoBot.

var defaultCriticalEnvironment = map[string]string{
	"GOROOT": "/usr/local/go",
	"GOPATH": Home() + "/go",
}
