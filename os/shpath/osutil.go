package shpath

import (
	"errors"
	"os"
	"strings"
)

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
