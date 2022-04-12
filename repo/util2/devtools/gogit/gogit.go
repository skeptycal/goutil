// Package gogit implements git cli commands in a more convenient way.
package gogit

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/skeptycal/goutil/repo/util2/gofile"
	"github.com/skeptycal/zsh"
)

const (
	gitCommitFormatString = `git commit -m '%s'`
)

// example: 6336b5a5ca051f416e63a8144eecf184cb1a3590
var isHash = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString

func IsHash(s string) bool {
	return isHash(s)
}

func IsAlphaNum(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func AddAll() error {
	return Err(zsh.Status("git add --all"))
}

func Add(s ...string) error {
	command := fmt.Sprintf("git add %s", strings.Join(s, " "))

	if err := Err(zsh.Status(command)); err != nil {
		return fmt.Errorf("error during command: %v", command)
	}
	return nil
}

// Commit creates a commit with message
func Commit(message string) error {
	command := fmt.Sprintf(gitCommitFormatString, message)
	return Err(zsh.Status(command))
}

// CommitAll creates a commit with message that
// contains all updated files.
func CommitAll(message string) error {
	if err := Err(AddAll()); err != nil {
		return err
	}
	return Err(Commit(message))
}

// GitInit initializes the Git environment in the current directory with:
//  git init
//  git add --all
//  git commit -m 'Initial Commit'
func GitInit() error {
	if err := Err(zsh.Status("git init")); err != nil {
		return err
	}
	return Err(CommitAll("Initial Commit"))
}

func PushTags() error {
	command := fmt.Sprintf("git push %s --tags", RemoteName())
	return Err(zsh.Status(command))
}

func getVersionCommitHash() string {
	return zsh.Sh("git rev-list --tags --max-count=1")
}
func VersionTag() string {
	return zsh.Sh("git describe --tags $(git rev-list --tags --max-count=1)")
}

// Tag create a git tag object signed with GPG
func Tag(s string) error {
	// todo check tag with regex
	if s == "" {
		return fmt.Errorf("git tag command invalid: %s", s)
	}
	args := strings.Split(s, " ")
	command := strings.TrimSpace(args[0])

	fmt.Printf("command: %s", command)

	tag := s[1:]
	return zsh.Status(fmt.Sprintf("git tag %s", tag))
}

// RemoteName gets the name of the remote branch, usually origin.
func RemoteName() string {
	return zsh.Sh("git remote")
}

// Remote returns the remote repository url.
/*
   origin	git@github.com:skeptycal/util.git (fetch)
   origin	git@github.com:skeptycal/util.git (push)
*/
func Remote() string {
	// todo - this is ... kinda messy
	remote := zsh.Sh("git remote")
	out := zsh.Sh("git remote -v")

	list := strings.Split(out, "\n")
	for _, s := range list {
		if strings.Contains(s, remote) {
			out = strings.TrimSpace(s)
			break
		}
	}
	s := ""
	for _, c := range out {
		if 32 < c && c < 127 {
			s += fmt.Sprintf("%c", c)
		} else {
			s += " "
		}
	}

	list = strings.Split(s, " ")
	// return fmt.Sprintf("list(%v): %v", len(list), list[1]) //! dev testing
	return list[1]
}

// Err calls error handling and logging routines
func Err(err error) error {
	return gofile.Err(err)
}
