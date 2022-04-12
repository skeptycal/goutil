package benchmark

import (
	"testing"

	"github.com/skeptycal/testes"
	"github.com/skeptycal/types"
)

var (
	NewAnyValue    = types.NewAnyValue
	AssertSameFunc = testes.AssertSameFunc
	AssertSameType = testes.AssertSameType
	AssertSameKind = testes.AssertSameKind
)

func TestNewCaller(t *testing.T) {
	type args struct {
		fn callerFunc
	}
	tests := []struct {
		name string
		args args
		want *caller
	}{
		// TODO: Add test cases.
		{"noop", args{fn: CallSetGlobalReturnValue}, &caller{fn: CallSetGlobalReturnValue, fnTrue: CallSetGlobalReturnValue, fnFalse: noop}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &caller{fn: tt.args.fn, fnTrue: tt.args.fn, fnFalse: noop}

			gFn := NewAnyValue(c.fn)
			wFn := NewAnyValue(tt.want.fn)

			AssertSameType(t, tt.name, gFn, wFn)
			AssertSameKind(t, tt.name, gFn, wFn)

			AssertSameFunc(t, tt.name, gFn, wFn)

		})
	}
}
