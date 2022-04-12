package generic

import "testing"

func TestStack(t *testing.T) {
	var x1 int = 123
	var y1 int = 456
	t.Run("int stack", func(t *testing.T) {
		myStackOfInts := new(Stack[int])
		AssertTrue(t, myStackOfInts.IsEmpty())  // check stack is empty
		myStackOfInts.Push(x1)                  // add a thing
		AssertFalse(t, myStackOfInts.IsEmpty()) // then check it's not empty
		myStackOfInts.Push(-y1)                 // add another thing
		value, _ := myStackOfInts.Pop()         // pop it back again
		AssertEqual(t, value, -y1, false)       // check value is equal
		value, _ = myStackOfInts.Pop()          // pop first value back
		AssertEqual(t, value, x1, false)        // check value is equal
		AssertTrue(t, myStackOfInts.IsEmpty())  // check stack is empty

		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		firstNum, _ := myStackOfInts.Pop()
		secondNum, _ := myStackOfInts.Pop()
		// we can add them because they are types, not interface{}
		AssertEqual(t, firstNum+secondNum, 3, false)
	})
	var u1 uint = 123
	var u2 uint = 456
	t.Run("uint stack", func(t *testing.T) {
		myStackOfUints := new(Stack[uint])
		AssertTrue(t, myStackOfUints.IsEmpty())  // check stack is empty
		myStackOfUints.Push(u1)                  // add a thing
		AssertFalse(t, myStackOfUints.IsEmpty()) // then check it's not empty
		myStackOfUints.Push(u2)                  // add another thing
		value, _ := myStackOfUints.Pop()         // pop it back again
		AssertEqual(t, value, u2, false)         // check value is equal
		value, _ = myStackOfUints.Pop()          // pop first value back
		AssertEqual(t, value, u1, false)         // check value is equal
		AssertTrue(t, myStackOfUints.IsEmpty())  // check stack is empty

		myStackOfUints.Push(1)
		myStackOfUints.Push(2)
		firstNum, _ := myStackOfUints.Pop()
		secondNum, _ := myStackOfUints.Pop()
		// we can add them because they are types, not interface{}
		AssertEqual(t, firstNum+secondNum, 3, false)
	})
	var s1 string = "123"
	var s2 string = "456"
	t.Run("uint stack", func(t *testing.T) {
		myStackOfUints := new(Stack[string])
		AssertTrue(t, myStackOfUints.IsEmpty())  // check stack is empty
		myStackOfUints.Push(s1)                  // add a thing
		AssertFalse(t, myStackOfUints.IsEmpty()) // then check it's not empty
		myStackOfUints.Push(s2)                  // add another thing
		value, _ := myStackOfUints.Pop()         // pop it back again
		AssertEqual(t, value, s2, false)         // check value is equal
		value, _ = myStackOfUints.Pop()          // pop first value back
		AssertEqual(t, value, s1, false)         // check value is equal
		AssertTrue(t, myStackOfUints.IsEmpty())  // check stack is empty

		myStackOfUints.Push("1")
		myStackOfUints.Push("2")
		firstNum, _ := myStackOfUints.Pop()
		secondNum, _ := myStackOfUints.Pop()
		// we can add them because they are types, not interface{}
		AssertEqual(t, firstNum+secondNum, "21", false)
	})
}
