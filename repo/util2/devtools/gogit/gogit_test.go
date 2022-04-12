// Package gogit implements git cli commands in a more convenient way.
package gogit

import (
	"testing"

	"
)

func TestRemote(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{"origin", "git@github.com:skeptycal/util.git"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remote(); got != tt.want {
				t.Errorf("Remote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoteName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{"origin", "origin"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoteName(); got != tt.want {
				t.Errorf("RemoteName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTag(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"empty command", args{""}, true},
		{"empty command", args{""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Tag(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("GitTag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPushTags(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"push current tags", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PushTags(); (err != nil) != tt.wantErr {
				t.Errorf("PushTags() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVersionTag(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		// {"current version", "v"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VersionTag(); got != tt.want {
				t.Errorf("VersionTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getVersionCommitHash(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"hash"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getVersionCommitHash()
			if !IsHash(got) {
				t.Errorf("getVersionCommitHash() = %v", got)
			}
		})
	}
}

func TestIsAlphaNum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"a1", args{"a1"}, true},
		{"%", args{"%"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlphaNum(tt.args.s); got != tt.want {
				t.Errorf("IsAlphaNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsHash(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"a1", args{"a1"}, true},
		{"%", args{"%"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHash(tt.args.s); got != tt.want {
				t.Errorf("IsHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddAll(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		// {"git add --all", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddAll(); (err != nil) != tt.wantErr {
				t.Errorf("AddAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	zsh.ShWait("mkdir -p tmp")
	zsh.ShWait("touch tmp/fake")
	zsh.ShWait("echo 'fake' >> tmp/fake")
	type args struct {
		s []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"nil", args{s: []string{""}}, true},
		{"add fake", args{s: []string{"tmp/fake"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Add(tt.args.s...); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	// zsh.Sh("rm -rf tmp/fake")
	// zsh.Sh("rm -rf tmp/")

}
