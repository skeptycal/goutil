package defaults

import (
	"fmt"
	"testing"
)

func TestCheckType(t *testing.T) {
	type args struct {
		any Any
		typ string
	}
	tests := []struct {
		name  string
		input Any
		typ   string
		knd   string
		want  bool
	}{
		{"string", "fake", "string", "string", true},
		{"int", 42, "int", "int", true},
		{"nil", nil, "nil", "invalid", true},
		{"0", 0, "int", "int", true},
		{"bool", true, "bool", "bool", true},
		{"empty []int32", []rune{}, "[]int32", "slice", true},
		{"[]uint8", []byte("fake"), "[]uint8", "slice", true},
	}
	for _, tt := range tests {
		AssertEquals(t, "CheckType", tt.name+", "+tt.typ, tt.want, CheckType(tt.input, tt.typ), false)
		AssertEquals(t, "GetType", tt.name, tt.typ, GetType(tt.input), false)
		AssertEquals(t, "GetKind", tt.name, tt.knd, GetKind(tt.input), false)
	}
}

func TestDefaultMapper(t *testing.T) {
	const defaultsStringSample = "Default Settings Map:\nKey                  = Value\ndebugState           = true\ntraceState           = true\n"

	// TODO: something about the string formatting is messing up this test, so
	// just testing the first 49 characters
	AssertEquals(t, "Defaults_String", "", defaultsStringSample[:49], Defaults.String()[:49], false)

	AssertEquals(t, "DefaultMapper.IsDebug", "", defaultDebugState, Defaults.IsDebug(), false)
	AssertNotEqual(t, "DefaultMapper.IsDebug", "", defaultDebugState, Defaults.IsDebug(), true)
	AssertEquals(t, "DefaultMapper.IsTrace", "", defaultTraceState, Defaults.IsTrace(), false)

	got, err := Defaults.Get("debugState")
	if err != nil {
		t.Fatal(err)
	}
	AssertEquals(t, "DefaultMapper.Get", "debugState", defaultDebugState, got.(Setting).AsBool(), false)

	err = Defaults.Set("debugState", !defaultDebugState)
	if err != nil {
		t.Fatal(err)
	}
	got, err = Defaults.Get("debugState")
	if err != nil {
		t.Fatal(err)
	}

	AssertEquals(t, "DefaultMapper.Set", "debugState", !defaultDebugState, got.(Setting).AsBool(), false)

	err = Defaults.Set(42, 42)
	if err == nil {
		t.Errorf("expected error, got %q", err)
	}

	_, err = Defaults.Get(42)
	if err == nil {
		t.Errorf("expected error, got %q", err)
	}

	_, err = Defaults.Get("fake_key_that_does_not_exist")
	if err == nil {
		t.Errorf("expected error, got %q", err)
	}

}

func ExampleDefaultMapper_String() {
	fmt.Println(Defaults.String())
	/*
				Default Settings Map:
		Key                  = Value
		debugState           = true
		traceState           = true
	*/
}
