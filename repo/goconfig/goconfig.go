package goconfig

import (
	"os"
	"os/exec"
	"strings"
)

var (
// c = Config{testUseGoRoutines: true}
)

type Config struct {
	homePath string
	goPath   string
	whoAmI   string
	path     string
	repoRoot string

	testConfig
}

type testConfig struct {
	testUseGoRoutines bool
}

func (c Config) RepoRoot() string {
	if c.repoRoot == "" {
		c.repoRoot = c.GoPath() + "/src/github.com/" + c.WhoAmI()
	}
	return c.repoRoot
}

func (c Config) WhoAmI() string {
	if c.whoAmI == "" {
		c.whoAmI = c.Sh("whoami")
	}
	return c.whoAmI
}

func (c Config) Path() string {
	if c.path == "" {
		c.path = c.Env("$PATH")
	}
	return c.path
}

func (c Config) GetHomeDir() string {
	if c.homePath == "" {
		c.homePath = c.Env("$HOME")
	}
	return c.homePath
}

func (Config) Env(s string) string     { return os.ExpandEnv(s) }
func (c Config) t_UseGoRoutines() bool { return c.testUseGoRoutines }

func (c Config) Sh(command string) string {
	b, err := c.Bytes(command)
	if err != nil {
		return err.Error()
	}
	return strings.TrimSpace(string(b))
}

func (c Config) Bytes(command string) ([]byte, error) {
	appargs := strings.Split(command, " ")
	cmd := exec.Command(appargs[0], appargs[1:]...)

	return cmd.Output()
}

func (c Config) GoPath() string { return c.Env("$GOPATH") }

// tes checks for the empty string
func tes(s string, got *string, want string) {
	if s == "" && *got != want {
		got = &want
	}
}
