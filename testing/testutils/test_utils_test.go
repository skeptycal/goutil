package testutils

import (
	"testing"
)

func TestAssertTypeEquals(t *testing.T) {
	if err, _ := IsType(42, "string"); err == nil {
		t.Error(err)
	}
	if err, _ := IsType(42, "int"); err != nil {
		t.Error(err)
	}
}

func TestAssertEquals(t *testing.T) {
	type args struct {
		testname string
		argname  string
		want     Any
		got      Any
		wantErr  bool
	}
	tests := []struct {
		name string
		args args
	}{
		{"equal ints", args{"fake", "42", 42, 42, false}},
		{"unequal ints", args{"fake", "42", 42, 24, true}},
		{"unequal strings", args{"AssertEquals", "string1", "string1", "string2", true}},
		{"equal strings", args{"AssertEquals", "string1", "string1", "string1", false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.args.wantErr {
				AssertEquals(t, tt.args.testname, tt.args.argname, tt.args.want, tt.args.got, tt.args.wantErr)
			} else {
				AssertNotEqual(t, tt.args.testname, tt.args.argname, tt.args.want, tt.args.got, !tt.args.wantErr)
			}
		})
	}
}

func TestIsString(t *testing.T) {
	tests := []struct {
		name    string
		arg     Any
		wantErr bool
	}{
		{"string", "string", false},
		{"nil", nil, true},
		{"int", 42, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err, _ := IsString(tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("IsString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAssertStringEquals(t *testing.T) {
	type args struct {
		testname string
		argname  string
		want     Any
		got      Any
		wantErr  bool
	}
	tests := []struct {
		name string
		args args
	}{
		{"equal ints", args{"AssertStringEquals", "42-42", 42, 42, true}},
		{"unequal ints", args{"AssertStringEquals", "42-24", 42, 24, true}},
		{"int string", args{"AssertStringEquals", "42-string1", 42, "string2", true}},
		{"string int", args{"AssertStringEquals", "string1-42", "string1", 42, true}},
		{"unequal strings", args{"AssertStringEquals", "string1-2", "string1", "string2", true}},
		{"equal strings", args{"AssertStringEquals", "equal strings", "string1", "string1", false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertStringEquals(t, tt.args.testname, tt.args.argname, tt.args.want, tt.args.got, tt.args.wantErr)
		})
	}
}
