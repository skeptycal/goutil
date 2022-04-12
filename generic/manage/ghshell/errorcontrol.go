package ghshell

// TODO try out code from VSCode extension:
//  Terminals Manager (currently v1.13.0) by Fabio Spampinato

import (
	"context"
	"math/rand"
	"time"
)

var ctx = context.TODO()

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	defaultOSTimeout = time.Second * 30
)

// TODO channel to close everything if a serious error occurs
// - wait channel
// - context ...
// - defer function to handle all function body errors
type errorControl struct {
	timeout time.Duration `default:"defaultOSTimeout"`
	cancel  chan struct{}
}

func NewErrorControl() *errorControl {
	return &errorControl{}
}

var ec = &errorControl{}

// TODO deferError handles all errors at the end of a
// function execution, including ...
// - logging, if active
// - checking and wrapping of errors
// - monitoring for Fatal errors
// - sending close signal to channel
func (e *errorControl) deferError(err *error) error {
	return *err
}

var de = ec.deferError
