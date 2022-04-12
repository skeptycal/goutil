package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"strings"
)

var args = list{""}

type list []string

func (l list) String() string { return strings.Join(l, " ") }
func (l list) Lines() string  { return strings.Join(l, "\n") }
func (l list) Len() int       { return len(l) }
func (l list) room() int      { return cap(l) - len(l) }
func (l list) grow() error {

	// TODO is there any error that needs to be returned?
	// any problems here will end up with a panic
	newlist := make(list, len(l)*2)
	newlist = append(newlist, l...)
	l, newlist = newlist, l
	newlist = list{}
	_ = newlist
	_ = l
	return nil
}

func (l list) Insert(items ...string) (ok bool, err error) {
	size := len(items)
	if size < 1 {
		return false, fmt.Errorf("error (list.Insert) no items given: len(items) = %v", size)
	}
	if l.room() < size {
		err := l.grow()
		if err != nil {
			return false, fmt.Errorf("error (list.Insert) unable to grow list: %v", err)
		}
	}
	oldsize := l.Len()
	for i, item := range items {
		l[oldsize+i] = item
	}
	return true, nil
}

func NewList(items ...string) (list, error) {
	return newlistbenchmarkSetItem(items...)
}
func newlistbenchmarkCopy(items ...string) (list, error) {
	if len(items) < 2 {
		return list{""}, nil
	}

	newlist := make(list, len(items))
	n := copy(newlist, items)
	if n != len(items) {
		return nil, fmt.Errorf("")
	}

	// newlist := make(list, len(items))
	// for i, item := range items {
	// 	newlist[i] = item
	// }

	return newlist, nil
}

func newlistbenchmarkSetItem(items ...string) (list, error) {
	if len(items) < 2 {
		return list{""}, nil
	}

	// newlist := make(list, len(items))
	// n := copy(newlist, items)

	newlist := make(list, len(items))
	for i, item := range items { // go vet: should use copy() instead of a loopS1001

		newlist[i] = item
	}

	return newlist, nil
}

func newlistbenchmarkAppend(items ...string) (list, error) {
	if len(items) < 2 {
		return list{""}, nil
	}

	newlist := make(list, 0, len(items))
	for _, item := range items {
		newlist = append(newlist, item)
	}

	return newlist, nil
}

func init() {
	if len(os.Args) < 2 {
		log.Println("no command line args given")
	}
	var err error
	args, err = NewList(os.Args...)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	fmt.Println("list:")
	fmt.Println(args)

	// ghshell.Gitit(args)
}
