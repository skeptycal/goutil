package types_test

// import (
// 	"fmt"
// 	"os"
// 	"sync"
// 	"testing"
// 	"time"
// )

// var setupDone bool = false
// var globalString string

// func SetupEnv(t *testing.T) {
// 	t.Helper()

// 	if !setupDone {
// 		t.Setenv("NO_COLOR", "true")
// 		setupDone = true
// 	}
// }

// func CleanupEnv(t *testing.T) {
// 	t.Helper()

// 	if setupDone {
// 		t.Cleanup(func() {
// 			os.Unsetenv("NO_COLOR")
// 		})
// 		setupDone = false
// 	}
// }

// func TestWithLock(t *testing.T) {
// 	lockfunc := func() {
// 		globalString = fmt.Sprintf("lock time: %v", time.Now())
// 	}
// 	type args struct {
// 		lk Locker
// 		fn func()
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		{"sync.Mutex", args{&sync.Mutex{}, lockfunc}},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			WithLock(tt.args.lk, tt.args.fn)
// 		})
// 	}
// }

// func TestContains(t *testing.T) {
// 	type args struct {
// 		needle   Any
// 		haystack []Any
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		{"true", args{"v", []Any{"v", "v"}}, true},
// 		{"wrong type", args{3, []Any{"v", "v"}}, false},
// 		{"wrong value", args{"d", []Any{"v", "v"}}, false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Contains(tt.args.needle, tt.args.haystack); got != tt.want {
// 				t.Errorf("Contains() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestCount(t *testing.T) {
// 	type args struct {
// 		needle   Any
// 		haystack []Any
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{"true", args{"v", []Any{"v", "v"}}, 2},
// 		{"true", args{"f", []Any{"v", "v", "f"}}, 1},
// 		{"wrong type", args{3, []Any{"v", "v"}}, 0},
// 		{"wrong value", args{"d", []Any{"v", "v"}}, 0},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Count(tt.args.needle, tt.args.haystack); got != tt.want {
// 				t.Errorf("Count() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// type fakeStringer struct {
// 	id   int
// 	name string
// }

// func (f *fakeStringer) String() string {
// 	return fmt.Sprintf("fakeid: %v,  name:%v", f.id, f.name)
// }

// func TestToString(t *testing.T) {
// 	// p := uint(42)

// 	tests := []struct {
// 		name string
// 		a    Any
// 		want string
// 	}{
// 		{"int 42", 42, "42"},
// 		{"uint 42", uint(42), "42"},
// 		// {"pointer uint 42", &p, "address"},
// 		{"Stringer", &fakeStringer{42, "42"}, "fakeid: 42,  name:42"},
// 		{"nil", nil, "<nil>"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := ToString(tt.a); got != tt.want {
// 				t.Errorf("ToString() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestIsComparable(t *testing.T) {
// 	for _, tt := range reflectTests {
// 		got := IsComparable(tt.a)
// 		k := ValueOf(tt.a).Kind()
// 		want := kindMaps[k].IsComparable()
// 		tRun(t, tt.name, got, want)
// 	}
// }

// func TestIsOrdered(t *testing.T) {
// 	for _, tt := range reflectTests {
// 		got := IsOrdered(tt.a)
// 		k := ValueOf(tt.a).Kind()
// 		want := kindMaps[k].IsOrdered()
// 		tRun(t, tt.name, got, want)
// 	}
// }

// func TestIsDeepComparable(t *testing.T) {
// 	for _, tt := range reflectTests {
// 		got := IsDeepComparable(tt.a)
// 		k := ValueOf(tt.a).Kind()
// 		want := kindMaps[k].IsDeepComparable()
// 		tRun(t, tt.name, got, want)
// 	}
// }

// func TestIsIterable(t *testing.T) {
// 	for _, tt := range reflectTests {
// 		got := IsIterable(tt.a)
// 		k := ValueOf(tt.a).Kind()
// 		want := kindMaps[k].IsIterable()
// 		tRun(t, tt.name, got, want)
// 	}
// }

// func TestHasAlternate(t *testing.T) {
// 	for _, tt := range reflectTests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := HasAlternate(tt.a)
// 			k := ValueOf(tt.a).Kind()
// 			want := kindMaps[k].HasAlternate()
// 			tRun(t, tt.name, got, want)
// 		})
// 	}
// }
