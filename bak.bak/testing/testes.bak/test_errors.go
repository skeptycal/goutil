package testes

import (
	"fmt"
	"testing"

	"github.com/skeptycal/goutil/types"
)

type ConfigOptions struct {
	OutputFieldLimit int // max size of data output field (default 0 = no limit)
}

var Config = NewConfig()

func NewConfig() ConfigOptions {
	return ConfigOptions{}
}

// var (
// 	LimitResult            bool // replaced with Config.OutputFieldLimit
// 	DefaultTestResultLimit = 15 // replaced with Config.OutputFieldLimit = 0
// )

func typeGuardExclude(needle Any, notAllowed []types.Any) bool {
	return !Contains(needle, notAllowed)
}

func TErrorf(t *testing.T, formatString assertFMT, name string, got, want Any) {
	if formatString == "" {
		formatString = "%v " + assertFmtSuffix // "%v = %v(%T), want %v(%T)"
	}
	t.Errorf(formatString.String(), name, limitTestResultLength(got), got, limitTestResultLength(want), want)
}

func TTypeError(t *testing.T, name string, got, want Any) {
	t.Errorf("%v = %v(%T), want %v(%T)", name, limitTestResultLength(got), got, limitTestResultLength(want), want)
}

func TError(t *testing.T, name string, got, want Any) {
	t.Errorf("%v = %v, want %v", name, limitTestResultLength(got), limitTestResultLength(want))
}

func limitTestResultLength(v Any) string {
	s := fmt.Sprintf("%v", v)

	if Config.OutputFieldLimit < 1 || len(s) < Config.OutputFieldLimit-3 {
		return s
	}

	return s[:Config.OutputFieldLimit-3] + "..."

}

func TName(testname, funcname, argname Any) string {
	if argname == "" {
		return fmt.Sprintf("%v: %v()", testname, funcname)
	}
	return fmt.Sprintf("%v: %v(%v)", testname, funcname, argname)
}
