package gotester

import (
	"os"

	errorlogger "
	"
)

type (
	Any = types.Any
)

var (
	log      = errorlogger.Log
	checkErr = errorlogger.Err
	pwd      string
)

func init() {
	var err error
	pwd, err = os.Getwd()
	Die(err)
}

// Check will log an error message if err is not nil
func Check(err error) error {
	if err == nil {
		return err
	}
	return checkErr(err)
}

// Die will exit with an error message if err is not nil
func Die(err error) {
	if err == nil {
		return
	}
	log.Fatal(err.Error())
}