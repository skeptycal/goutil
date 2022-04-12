// Package polynomial provides functions that support polynomial arithmetic.
package polynomial

import (
	"fmt"
	"strings"
)

// New returns a new List with each digit of the value n
// set to its own ListNode.
func New(n int) *List {

    current := &ListNode{}
    list := &List{false, current, nil}

    if n < 0 {
        n = -n
        list.SignBit = true
    }

	s := fmt.Sprintf("%d", n)


	for _, r := range s {
		current.Val = int(r - 48)
		current.Next = &ListNode{}
		current = current.Next
    }

	return list
}

// ListNode defines a singly-linked list of integer values.
type ListNode struct {
	Val  int       `default:"0"`
	Next *ListNode `default:"nil"`
}

// List defines the boundaries of a list signly linked list.
type List struct {
	SignBit bool      `default:"false"`
	First   *ListNode `default:"&ListNode{}"`
	Last    *ListNode `default:"nil"`
}

func (l *List) Free() {
    l = &List{}
}

// LoadInt loads the digits of n into a List and returns
// the String() response. The digits will be in reverse order.
func (l *List) LoadInt(n int) string {
    l = ListDigits(n)
    return l.String()
}

func ListDigits(n int)  *List {

    current := &ListNode{}
    list := &List{false, current, nil}

	if n == 0 {
		return &List{false, &ListNode{0,nil}, nil}
	}

	if n < 0 {
		list.SignBit = true
		n = -n
	}



    current = list.First
    tmp := 0
	for n > 0 {
        tmp = n % 10
        current.Val = tmp
        current.Next = &ListNode{}
        current = current.Next
		n /= 10
	}
	return list
}

func (l *List) String() string {

    if l.First.Val == 0 {
        return "0"
    }

    sb := strings.Builder{}
    if l.SignBit {
        sb.WriteString("-")
    }

    current := l.First

	for current.Next != nil {
        sb.WriteRune(rune(current.Val+48))
		current = current.Next
    }

	return sb.String()
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//     var num int
//     num += l1.Val + l2.Val

//     for {
//         n1 := l1.Val
//     }
// }

// StringDigits parses an integer value one digit at a time, starting
// with least significant, and returns a string representation.
func StringDigits(n int) string {

	if n == 0 {
		return "0"
	}

	result := strings.Builder{}
	defer result.Reset()

	if n < 0 {
		result.WriteString("-")
		n = -n
	}

	tmp := 0
	for n > 0 {
		tmp = n % 10
		result.WriteByte(byte(tmp + 48))
		n /= 10
	}
	return result.String()
}
