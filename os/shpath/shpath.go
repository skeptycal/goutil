package shpath

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	NL        = "\n"
	nlWindows = "\r\n"
	TAB       = "\t"
	PATHSEP   = string(os.PathListSeparator)
)

var Verbose bool = false

type ShPath struct{ list []string }

func NewPath() *ShPath {
	p := &ShPath{}
	err := p.load()
	if err != nil {
		return nil
	}
	return p
}

// Clean removes invalid directories,
// and returns the number removed, if any.
// The order of the directories is maintained.
func (p *ShPath) Clean() (n int) {
	for i, v := range p.list {
		if v == "" || !IsDir(v) {
			p.list = RemoveOrdered(p.list, i)
			n += 1
			if Verbose {
				fmt.Fprintf(os.Stderr, "the path (%v) is not a valid directory\n", v)
			}
		}
	}

	if n > 0 && Verbose {
		fmt.Fprintf(os.Stderr, "directories checked (%v removed)\n", n)
	}

	return n
}

func (p *ShPath) load() error {
	s, err := GetEnvValue("path")
	if err != nil {
		return err
	}
	s = DropDupes(s, PATHSEP)
	s = strings.Replace(s, nlWindows, NL, -1)
	p.list = strings.Split(s, PATHSEP)
	return nil
}

// Out returns the path in delimited format ready
// for OS use. Out runs Clean() on the list.
func (p *ShPath) Out() string {
	p.Clean()
	return strings.Join(p.list, PATHSEP)
}

// Add checks that the directory exists and
// adds element s to the path at position n.
func (p *ShPath) Add(s string, n int) error {
	if s == "" {
		return errors.New("path cannot be empty")
	}
	if !IsDir(s) {
		v := fmt.Sprintf("the path (%v) is not a valid directory\n", s)
		if Verbose {
			fmt.Fprint(os.Stderr, v)
		}
		return errors.New(v)
	}

	// if n is out of bounds, append
	// s to end of list
	if n < 0 || n >= len(p.list) {
		p.list = Append(p.list, s)
		return nil
	}

	p.list = Insert(p.list, s, n)

	return nil
}

func (p *ShPath) Len() int {
	return len(p.list)
}

// String returns the path in newline delimited format.
func (p *ShPath) String() string {
	return strings.Join(p.list, NL)
}

func (p *ShPath) DebugPrint() {
	fmt.Println("path.DebugPrint()")
	for i, v := range p.list {
		fmt.Printf("%3d: %s\n", i, v)
	}
}
