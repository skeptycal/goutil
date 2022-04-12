// Copyright (c) 2021 Michael Treanor
// https://
// MIT License

package errorlogger

import (
	"errors"
	"os"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	for _, tt := range errorloggerTests {
		t.Run(tt.name, func(t *testing.T) {

			got := newTestStruct(tt.args.enabled, tt.args.msg, tt.args.wrap, tt.args.fn, tt.args.logger)

			got.SetLoggerFunc(nil)
			got.SetLoggerFunc(got.Debug)
			err := got.SetLogLevel("")
			if err == nil {
				t.Errorf("New(%s) setting log level to the empty string should produce an error: %v", tt.name, err)
			}
			err = got.SetLogLevel("INFO")
			if err != nil {
				t.Errorf("New(%s) setting log level correctly should not produce an error: %v", tt.name, err)
			}

			err = got.SetLogOutput(nil)
			if err == nil {
				t.Errorf("New(%s) setting log output to nil should produce an error: %v", tt.name, err)
			}
			err = got.SetLogOutput(os.Stdout)
			if err != nil {
				t.Errorf("New(%s) setting log output correctly should not produce an error: %v", tt.name, err)
			}

			got.SetJSON(true)
			got.SetText()

			switch tt.want.(type) {
			case ErrorLogger:
				if tt.wantErr {
					t.Errorf("New(%s) implements ErrorLogger: got %T, want %T", tt.name, got, tt.want)
				}
			default:
				if !tt.wantErr {
					t.Errorf("New(%s) does not implement ErrorLogger: got %T, want %T", tt.name, got, tt.want)
				}
			}
		})
	}
}

func Test_errorLogger_SetErrorWrap(t *testing.T) {
	tests := []struct {
		name  string
		input error
		wrap  error
	}{
		{"fakeError", errFake, fakeSysCallError},
		{"nil", nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := NewWithOptions(true, "", nil, nil, nil)
			got.SetErrorWrap(tt.wrap)

			if errors.Is(got.Err(errFake), fakeSysCallError) {
				t.Errorf("SetErrorWrap(%s) did not wrap error: got %v, want %v", tt.name, got, tt.wrap)
			}
		})
	}
}

func Test_errorLogger_SetCustomMessage(t *testing.T) {
	tests := []struct {
		name  string
		input error
		msg   string
	}{
		{"fakeError", errFake, "fakeMessage"},
		{"nil", nil, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := NewWithOptions(true, "", nil, nil, nil)
			got.SetCustomMessage(tt.msg)

			if tt.msg != "" && strings.Contains(got.Err(tt.input).Error(), tt.msg) {
				t.Errorf("SetErrorWrap(%s) did not wrap error: got %v, want %v", tt.name, got, tt.msg)
			}
		})
	}
}
