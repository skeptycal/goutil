package shpath

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	NL      = "\n"
	ListSep = ":"
)

type ShPath []string

func (p ShPath) Load() error {

	s, err := GetEnvValue("path")
	if err != nil {
		return err
	}
	
	for _, v := range strings.Split(s, ListSep) {
		if IsDir(v) {
			p = append(p, v)
		} else {
			fmt.Fprintf(os.Stderr, "the path (%v) is not a valid directory\n", v)
		}
	}

	return nil
}

func (p ShPath) Save(path string) error {

	s := strings.Join(p, ListSep)

	return os.Setenv("PATH", s)
}

func (p ShPath) String() string {
	sb := new(strings.Builder)
	defer sb.Reset()

	for _, v := range p {
		fmt.Fprintf(sb, " %s\n", v)
	}
	return sb.String()
}

func IsDir(s string) bool {
	fi, err := os.Stat(s)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func IsRegular(s string) bool {
	fi, err := os.Stat(s)
	if err != nil {
		return false
	}
	return fi.Mode().IsRegular()
}

func GetEnvValue(key string) (string, error) {
	key = strings.ToUpper(key)
	s, ok := os.LookupEnv(key)
	if !ok {
		return "", errors.New("environment variable (" + key + ") not found")
	}

	return s, nil
}
