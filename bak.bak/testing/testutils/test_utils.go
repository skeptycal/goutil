package testutils

import (
	"fmt"
	"testing"
)

func AssertEquals(t *testing.T, testname string, argname string, want Any, got Any, wantErr bool) {
	name := testname + "(" + argname + ")"
	t.Run(name, func(t *testing.T) {
		if got != want != wantErr {
			t.Errorf("%v = %v, want %v", name, got, want)
		}
	})
}

func AssertNotEqual(t *testing.T, testname string, argname string, want Any, got Any, wantErr bool) {
	name := testname + "(" + argname + ")"
	t.Run(name, func(t *testing.T) {
		if got == want != wantErr {
			t.Errorf("%v = %v, want %v", name, got, want)
		}
	})
}

func AssertStringEquals(t *testing.T, testname string, argname string, want Any, got Any, wantErr bool) {

	if err, _ := IsString(want); err != nil {
		if !wantErr {
			t.Errorf("%v(%v): %v", testname, argname, err)
			t.FailNow()
		}
		return
	}

	if err, _ := IsString(got); err != nil {
		if !wantErr {
			t.Errorf("%v(%v): %v", testname, argname, err)
			t.FailNow()
		}
		return
	}

	w := want.(string)
	g := got.(string)

	length := len(w)
	if len(g) < len(w) {
		length = len(g)
	}

	for i := 0; i < length; i++ {
		if w[i] != g[i] {
			err := fmt.Errorf("%v(%v): first string mismatch at position %d - want: %q  got: %q", testname, argname, i, w[i], g[i])
			if !wantErr {
				t.Error(err)
			} else {
				t.Log(err)
			}
			break
		}
	}
}

func IsType(arg Any, typ string) (err error, ok bool) {
	if a := GetType(arg); a != typ {
		return fmt.Errorf("incorrect type: %v want %v", a, typ), false
	}
	return nil, true
}

func IsString(arg Any) (err error, ok bool) { return IsType(arg, "string") }
