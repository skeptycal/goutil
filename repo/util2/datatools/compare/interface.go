package compare

import (
	log "github.com/sirupsen/logrus"
)

// InterfaceEqual protects against panics from doing equality tests on
// two interfaces with non-comparable underlying types.
// adapted from:
//
// /usr/local/go/src/os/exec/exec.go (go 1.15.6)
func InterfaceEqual(a, b interface{}) bool {
	defer func() {
		err := recover()
		log.Errorf("panic recovered: %v", err)
	}()
	return a == b
}
