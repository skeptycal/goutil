package errorlogger

import (
	"sync"
	. "sync"
	"sync/atomic"

	// . "sync/atomic"
	"testing"
	// . "testing"
	"fmt"
	"math/rand"
	"time"
)

func TestSize1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.got(tt.in); got != tt.want {
				t.Errorf("Size(%v) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

type (
	Test struct {
		name    string
		in      int
		got     func(int) string
		want    string
		wantErr bool
	}

	result bool
)

var tests = []Test{
	{"negative", -1, size, "negative", false},
	{"small", 5, size, "small", false},
	// {9, Size, "", true},
}

func results(t *testing.T, tests ...Test) <-chan result {
	c := make(chan result)
	go func() {
		for _, tt := range tests {
			got := tt.got(tt.in)
			if got == tt.want { // != tt.wantErr {
				t.Errorf("%sf(%v)=%v; want %v", tt.name, tt.in, got, tt.want)
				c <- true
			}
			c <- true
		}
	}()
	return c
}

func TestSize(t *testing.T) {

	// c := boring("boring!")
	c := results(t, tests...)

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %v\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")

	for _, test := range tests {
		go func() {
			RunTest(t, test)
		}()
	}
}

func RunTest(t *testing.T, tt Test) {
	got := tt.got(tt.in)
	if got == tt.want { // != tt.wantErr {
		t.Errorf("%sf(%v)=%v; want %v", tt.name, tt.in, got, tt.want)
	}
}

func TestWG(t *testing.T) {
	wg1 := new(sync.WaitGroup)
	wg2 := &WaitGroup{}

	// Run the same test a few times to ensure barrier is in a proper state.
	for i := 0; i != 8; i++ {
		testWaitGroup(t, wg1, wg2)
	}
}

/////////////////////////////// Rob Pike talk ////////////////////////////////
/////////////////////////////// Go Concurrency Patterns //////////////////////
/////////////////////////////// Google I/O 2012 //////////////////////////////
func TestBoring1(t *testing.T) {
	t.Log("Boring:")
	c := boring("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}

// boring is boring
// Reference: youtube talk by Rob Pike
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// Reference: go standard library sync package waitgroup.go ...

/*
BenchmarkWaitGroupUncontended-8    	   350057320	 	3.837 ns/op	       0 B/op	       0 allocs/op
BenchmarkWaitGroupAddDone-8        		10891663		101.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkWaitGroupAddDoneWork-8    		10495297		111.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkWaitGroupWait-8           	  1000000000	   0.7290 ns/op	       0 B/op	       0 allocs/op
BenchmarkWaitGroupWaitWork-8       	   136401049	 	9.443 ns/op	       0 B/op	       0 allocs/op
BenchmarkWaitGroupActuallyWait-8   		10656708		125.9 ns/op	      32 B/op	       2 allocs/op
*/

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package sync_test

// import (
// 	. "sync"
// 	"sync/atomic"
// 	"testing"
// )

func testWaitGroup(t *testing.T, wg1 *WaitGroup, wg2 *WaitGroup) {
	n := 16
	wg1.Add(n)
	wg2.Add(n)
	exited := make(chan bool, n)
	for i := 0; i != n; i++ {
		go func() {
			wg1.Done()
			wg2.Wait()
			exited <- true
		}()
	}
	wg1.Wait()
	for i := 0; i != n; i++ {
		select {
		case <-exited:
			t.Fatal("WaitGroup released group too soon")
		default:
		}
		wg2.Done()
	}
	for i := 0; i != n; i++ {
		<-exited // Will block if barrier fails to unlock someone.
	}
}

func TestWaitGroup(t *testing.T) {
	wg1 := &WaitGroup{}
	wg2 := &WaitGroup{}

	// Run the same test a few times to ensure barrier is in a proper state.
	for i := 0; i != 8; i++ {
		testWaitGroup(t, wg1, wg2)
	}
}

func TestWaitGroupMisuse(t *testing.T) {
	defer func() {
		err := recover()
		if err != "sync: negative WaitGroup counter" {
			t.Fatalf("Unexpected panic: %#v", err)
		}
	}()
	wg := &WaitGroup{}
	wg.Add(1)
	wg.Done()
	wg.Done()
	t.Fatal("Should panic")
}

func TestWaitGroupRace(t *testing.T) {
	// Run this test for about 1ms.
	for i := 0; i < 1000; i++ {
		wg := &WaitGroup{}
		n := new(int32)
		// spawn goroutine 1
		wg.Add(1)
		go func() {
			atomic.AddInt32(n, 1)
			wg.Done()
		}()
		// spawn goroutine 2
		wg.Add(1)
		go func() {
			atomic.AddInt32(n, 1)
			wg.Done()
		}()
		// Wait for goroutine 1 and 2
		wg.Wait()
		if atomic.LoadInt32(n) != 2 {
			t.Fatal("Spurious wakeup from Wait")
		}
	}
}

func TestWaitGroupAlign(t *testing.T) {
	type X struct {
		x  byte
		wg WaitGroup
	}
	var x X
	x.wg.Add(1)
	go func(x *X) {
		x.wg.Done()
	}(&x)
	x.wg.Wait()
}

func BenchmarkWaitGroupUncontended(b *testing.B) {
	type PaddedWaitGroup struct {
		WaitGroup
		pad [128]uint8
	}
	b.RunParallel(func(pb *testing.PB) {
		var wg PaddedWaitGroup
		for pb.Next() {
			wg.Add(1)
			wg.Done()
			wg.Wait()
		}
	})
}

func benchmarkWaitGroupAddDone(b *testing.B, localWork int) {
	var wg WaitGroup
	b.RunParallel(func(pb *testing.PB) {
		foo := 0
		for pb.Next() {
			wg.Add(1)
			for i := 0; i < localWork; i++ {
				foo *= 2
				foo /= 2
			}
			wg.Done()
		}
		_ = foo
	})
}

func BenchmarkWaitGroupAddDone(b *testing.B) {
	benchmarkWaitGroupAddDone(b, 0)
}

func BenchmarkWaitGroupAddDoneWork(b *testing.B) {
	benchmarkWaitGroupAddDone(b, 100)
}

func benchmarkWaitGroupWait(b *testing.B, localWork int) {
	var wg WaitGroup
	b.RunParallel(func(pb *testing.PB) {
		foo := 0
		for pb.Next() {
			wg.Wait()
			for i := 0; i < localWork; i++ {
				foo *= 2
				foo /= 2
			}
		}
		_ = foo
	})
}

func BenchmarkWaitGroupWait(b *testing.B) {
	benchmarkWaitGroupWait(b, 0)
}

func BenchmarkWaitGroupWaitWork(b *testing.B) {
	benchmarkWaitGroupWait(b, 100)
}

func BenchmarkWaitGroupActuallyWait(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var wg WaitGroup
			wg.Add(1)
			go func() {
				wg.Done()
			}()
			wg.Wait()
		}
	})
}
