package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"reflect"
	"strings"
	"testing"
	"time"
)

/*//* Using pprof:

## pprof basics
>pprof lets you collect CPU profiles, traces, and heap profiles for your Go programs. The normal way to use pprof seems to be:

- Set up a webserver for getting Go profiles (with `import _ "net/http/pprof"`)
- Run `curl localhost:$PORT/debug/pprof/$PROFILE_TYPE` to save a profile
- Use `go tool pprof` to analyze said profile
- You can also generate pprof profiles in your code using the pprof package but I haven’t done that.

## What’s a profile? What kinds of profiles can I get?

>A Profile is a collection of stack traces showing the call sequences that led to instances of a particular event, such as allocation. Packages can create and maintain their own profiles; the most common use is for tracking resources that must be explicitly closed, such as files or network connections.

Each Profile has a unique name. A few profiles are predefined:
goroutine    - stack traces of all current goroutines
heap         - a sampling of all heap allocations
threadcreate - stack traces that led to the creation of new OS threads
block        - stack traces that led to blocking on synchronization primitives
mutex        - stack traces of holders of contended mutexes

There are 7 places you can get profiles in the default webserver: the ones mentioned above

- [goroutine](http://localhost:6060/debug/pprof/goroutine)
- [heap](http://localhost:6060/debug/pprof/heap)
- [threadcreate](http://localhost:6060/debug/pprof/threadcreate)
- [block](http://localhost:6060/debug/pprof/block)
- [mutex](http://localhost:6060/debug/pprof/mutex)

and also 2 more: the CPU profile and the CPU trace.

- [profile](http://localhost:6060/debug/pprof/profile)
- [trace](http://localhost:6060/debug/pprof/trace?seconds=5)

To analyze these profiles (lists of stack traces), the tool to use is go tool pprof, which is a bunch of tools for visualizing stack traces.

**super confusing note**: the trace endpoint (`/debug/pprof/trace?seconds=5`), unlike all the rest, outputs a file that is not a pprof profile. Instead it’s a trace and you can view it using go tool trace (not go tool pprof).

You can see the available profiles with http://localhost:6060/debug/pprof/ in your browser. Except it doesn’t tell you about `/debug/pprof/profile` or `/debug/pprof/trace` for some reason.

All of these kinds of profiles (goroutine, heap allocations, etc) are just collections of stacktraces, maybe with some metadata attached. If we look at the pprof protobuf definition, you see that a profile is mostly a bunch of Samples.

A sample is basically a stack trace. That stack trace might have some extra information attached to it! For example in a heap profile, the stack trace has a number of bytes of memory attached to it. I think the Samples are the most important part of the profile.

We’re going to deconstruct what **exactly** is inside a pprof file later, but for now let’s start by doing a quick example of what analyzing a heap profile looks like!

## Example:

```go
$ go tool pprof  http://localhost:6060/debug/pprof/heap
    Fetching profile from http://localhost:6060/debug/pprof/heap
    Saved profile in /home/bork/pprof/pprof.localhost:6060.inuse_objects.inuse_space.004.pb.gz
    Entering interactive mode (type "help" for commands)
(pprof) top
    34416.04kB of 34416.04kB total (  100%)
    Showing top 10 nodes out of 16 (cum >= 512.04kB)
          flat  flat%   sum%        cum   cum%
       33904kB 98.51% 98.51%    33904kB 98.51%  main.leakyFunction
```

I can also do the same thing outside interactive mode with `go tool pprof -top http://localhost:6060/debug/pprof/heap`.

This basically tells us that `main.leakyFunction` is using 33.9MB of memory. Neat!

We can also generate a PNG profile like this: `go tool pprof -png http://localhost:6060/debug/pprof/heap > out.png`.


Reference: https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/
*/

type (
	testItem struct {
		name    string
		items   []string
		want    list
		wantErr bool
	}

	testSet = []testItem
)

func init() {
	rand.Seed(time.Now().UnixNano())
	runPProfServer()
}

func runPProfServer() error {

	var err error

	// we need a webserver to get the pprof webserver
	go func() {
		err := http.ListenAndServe("localhost:6060", nil)
		if err != nil {
			log.Println(err)

		}
	}()
	fmt.Println("pprof started")
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go BenchmarkNewlistBenchmarksAll()
	// wg.Wait()
	return err
}

func randString(n int) string {
	b := make([]byte, n)
	nn, err := rand.Read(b)
	if err != nil || nn != n {
		log.Fatalf("error reading random bytes: want %v, got %v", n, nn)
	}
	return string(b)
}

func randStringList(n int) []string {
	newlist := make([]string, n)
	for i := 0; i < n; i++ {
		newlist[i] = randString(rand.Intn(10) + 8)
	}
	return newlist
}

var randList = randStringList(2 << 8)

var tests = testSet{
	{"empty slice", []string{}, list{}, false},
	{"empty string", []string{}, list{""}, false},
	{"empty list", []string{}, list{}, false},
	{"nil", nil, list{}, false},
	{"abc", []string{"a", "b", "c"}, list{"a", "b", "c"}, false},
	{"long", randList, randList, false},
}

/*
//* Summary: copy is the way to Go in version 1.18+

/newlistbenchmarkCopy:_abc-8         	24708834	        47.35 ns/op	      48 B/op	       1 allocs/op
/newlistbenchmarkSetItem:_abc-8      	24161194	        49.15 ns/op	      48 B/op	       1 allocs/op
/newlistbenchmarkAppend:_abc-8       	24135742	        49.79 ns/op	      48 B/op	       1 allocs/op
/newlistbenchmarkCopy:_long-8        	  947654	      1323 ns/op	    8192 B/op	       1 allocs/op
/newlistbenchmarkSetItem:_long-8     	  499315	      2388 ns/op	    8192 B/op	       1 allocs/op
/newlistbenchmarkAppend:_long-8      	  495697	      2393 ns/op	    8192 B/op	       1 allocs/op

* More at the end of the file...

/newlistbenchmark1:_abc-8  						       	25258771	  	    47.25 ns/op	      48 B/op	       1 allocs/op
/newlistbenchmark2:_abc-8  						       	24260864	  	    49.23 ns/op	      48 B/op	       1 allocs/op
/newlistbenchmark1:_long-8 						       	  920012	  	     1343 ns/op	    8192 B/op	       1 allocs/op
/newlistbenchmark2:_long-8 						       	  496264	  	     2407 ns/op	    8192 B/op	       1 allocs/op
/newlistbenchmark1:_random_string:_len(4)-8         	24623696	        48.00 ns/op	      64 B/op	       1 allocs/op
/newlistbenchmark2:_random_string:_len(4)-8         	23703937	        50.42 ns/op	      64 B/op	       1 allocs/op
/newlistbenchmark1:_random_string:_len(8)-8         	21180426	        56.94 ns/op	     128 B/op	       1 allocs/op
/newlistbenchmark2:_random_string:_len(8)-8         	18592940	        64.61 ns/op	     128 B/op	       1 allocs/op
/newlistbenchmark1:_random_string:_len(16)-8        	15671174	        77.57 ns/op	     256 B/op	       1 allocs/op
/newlistbenchmark2:_random_string:_len(16)-8        	12712560	        94.65 ns/op	     256 B/op	       1 allocs/op
/newlistbenchmark1:_random_string:_len(32)-8        	10223976	        110.7 ns/op	     512 B/op	       1 allocs/op
/newlistbenchmark2:_random_string:_len(32)-8        	 8037162	        149.8 ns/op	     512 B/op	       1 allocs/op
/newlistbenchmark1:_random_string:_len(64)-8        	 6418945	        190.4 ns/op	    1024 B/op	       1 allocs/op
/newlistbenchmark2:_random_string:_len(64)-8        	 4333555	        277.8 ns/op	    1024 B/op	       1 allocs/op
/newlistbenchmark1:_random_string:_len(128)-8       	 3360402	        354.1 ns/op	    2048 B/op	       1 allocs/op
/newlistbenchmark2:_random_string:_len(128)-8       	 2441515	        494.4 ns/op	    2048 B/op	       1 allocs/op
/newlistbenchmark1:_random_string:_len(256)-8       	 1761536	        681.1 ns/op	    4096 B/op	       1 allocs/op
/newlistbenchmark2:_random_string:_len(256)-8       	 1000000	         1082 ns/op	    4096 B/op	       1 allocs/op


/// All nil or empty benchmarks are similar ... no further benchmarks
/newlistbenchmark1:_empty_string-8         	45917563	        25.76 ns/op	      16 B/op	       1 allocs/op
/newlistbenchmark2:_empty_string-8         	47199806	        25.14 ns/op	      16 B/op	       1 allocs/op
/newlistbenchmark1:_empty_list-8           	46069165	        25.73 ns/op	      16 B/op	       1 allocs/op
/newlistbenchmark2:_empty_list-8           	46914783	        25.13 ns/op	      16 B/op	       1 allocs/op
/newlistbenchmark1:_nil-8                  	45777356	        25.73 ns/op	      16 B/op	       1 allocs/op
/newlistbenchmark2:_nil-8                  	47187507	        25.12 ns/op	      16 B/op	       1 allocs/op
*/
func BenchmarkNewlistBenchmarksAll(b *testing.B) {

	funclist := []struct {
		name string
		fn   func(items ...string) (list, error)
	}{
		{"newlistbenchmarkCopy", newlistbenchmarkCopy},
		{"newlistbenchmarkSetItem", newlistbenchmarkSetItem},
		{"newlistbenchmarkAppend", newlistbenchmarkAppend},
	}
	var benchmarks testSet = append(testSet{}, tests[3:]...)
	for i := 1; i < 8; i++ {
		// newtest := 	{"long", randList, randList, false},
		size := 2 << i
		iName := fmt.Sprintf("random string: len(%v)", size)
		iStr := randStringList(size)
		benchmarks = append(benchmarks, testSet{
			testItem{iName, iStr, iStr, false},
		}...)

	}
	for _, bb := range benchmarks {

		for _, ff := range funclist {

			b.Run(ff.name+": "+bb.name, func(b *testing.B) {

				for i := 0; i < b.N; i++ {
					var retval any = []string{}
					switch len(bb.items) {
					case 0:
						retval = []string{}
					case 1:
						retval = []string{bb.items[0]}
					default:
						var err error
						retval, err = ff.fn(bb.items...)
						if err != nil {
							_ = errors.New("benchmark error")
							b.Errorf("error during benchmark: %v", err)
						}
					}
					_ = retval
				}
			})
		}
	}

}

func Test_newlistbenchmark1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newlistbenchmarkCopy(tt.items...)
			if (err != nil) != tt.wantErr {
				t.Errorf("newlistbenchmark1(%v) error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			// uses listString to compare space separated strings of list items ... which is the real goal here
			if listString(got) != listString(tt.want) && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newlistbenchmark1(%v) = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}

func listString(l list) string {
	return strings.Join(l, " ")
}
func Test_newlistbenchmark2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newlistbenchmarkSetItem(tt.items...)
			if (err != nil) != tt.wantErr {
				t.Errorf("newlistbenchmark2(%v) error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			// uses listString to compare space separated strings of list items ... which is the real goal here
			if listString(got) != listString(tt.want) && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newlistbenchmark2(%v) = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}

/*

* After adding in benchmarks with increasing list size (random strings)

/newlistbenchmark1:_abc-8  						       	25258771	  	    47.25 ns/op	      48 B/op	       1 allocs/op
/newlistbenchmark2:_abc-8  						       	24260864	  	    49.23 ns/op	      48 B/op	       1 allocs/op
/newlistbenchmark1:_long-8 						       	  920012	  	     1343 ns/op	    8192 B/op	       1 allocs/op
/newlistbenchmark2:_long-8 						       	  496264	  	     2407 ns/op	    8192 B/op	       1 allocs/op
/newlistbenchmark1:_random_string:_len(4)-8         	24623696	        48.00 ns/op	      64 B/op	       1 allocs/op
/newlistbenchmark2:_random_string:_len(4)-8         	23703937	        50.42 ns/op	      64 B/op	       1 allocs/op
/newlistbenchmark1:_random_string:_len(8)-8         	21180426	        56.94 ns/op	     128 B/op	       1 allocs/op
/newlistbenchmark2:_random_string:_len(8)-8         	18592940	        64.61 ns/op	     128 B/op	       1 allocs/op
/newlistbenchmark1:_random_string:_len(16)-8        	15671174	        77.57 ns/op	     256 B/op	       1 allocs/op
/newlistbenchmark2:_random_string:_len(16)-8        	12712560	        94.65 ns/op	     256 B/op	       1 allocs/op
/newlistbenchmark1:_random_string:_len(32)-8        	10223976	        110.7 ns/op	     512 B/op	       1 allocs/op
/newlistbenchmark2:_random_string:_len(32)-8        	 8037162	        149.8 ns/op	     512 B/op	       1 allocs/op
/newlistbenchmark1:_random_string:_len(64)-8        	 6418945	        190.4 ns/op	    1024 B/op	       1 allocs/op
/newlistbenchmark2:_random_string:_len(64)-8        	 4333555	        277.8 ns/op	    1024 B/op	       1 allocs/op
/newlistbenchmark1:_random_string:_len(128)-8       	 3360402	        354.1 ns/op	    2048 B/op	       1 allocs/op
/newlistbenchmark2:_random_string:_len(128)-8       	 2441515	        494.4 ns/op	    2048 B/op	       1 allocs/op
/newlistbenchmark1:_random_string:_len(256)-8       	 1761536	        681.1 ns/op	    4096 B/op	       1 allocs/op
/newlistbenchmark2:_random_string:_len(256)-8       	 1000000	         1082 ns/op	    4096 B/op	       1 allocs/op


//* After breakout of checks for list of length 0 and 1 to a common caller:
//* 	This resulted in a 60-70X increase in checks for empty lists ... very nearly 2 orders of magnitude ...
//*  	with no change in other operations.

/newlistbenchmarkCopy:_nil-8         	1000000000	         0.6843 ns/op	       0 B/op	       0 allocs/op
/newlistbenchmarkSetItem:_nil-8      	1000000000	         0.8154 ns/op	       0 B/op	       0 allocs/op
/newlistbenchmarkAppend:_nil-8       	1000000000	         0.6865 ns/op	       0 B/op	       0 allocs/op
/newlistbenchmarkCopy:_abc-8         	24202778	        48.98 ns/op	      48 B/op	       1 allocs/op
/newlistbenchmarkSetItem:_abc-8      	23771712	        50.37 ns/op	      48 B/op	       1 allocs/op
/newlistbenchmarkAppend:_abc-8       	23191485	        51.48 ns/op	      48 B/op	       1 allocs/op
/newlistbenchmarkCopy:_long-8        	  926175	      1272 ns/op	    8192 B/op	       1 allocs/op
/newlistbenchmarkSetItem:_long-8     	  503470	      2390 ns/op	    8192 B/op	       1 allocs/op
/newlistbenchmarkAppend:_long-8      	  504675	      2395 ns/op	    8192 B/op	       1 allocs/op
/newlistbenchmarkCopy:_random_string:_len(4)-8         	24647678	        48.60 ns/op	      64 B/op	       1 allocs/op
/newlistbenchmarkSetItem:_random_string:_len(4)-8      	23330775	        51.38 ns/op	      64 B/op	       1 allocs/op
/newlistbenchmarkAppend:_random_string:_len(4)-8       	22524212	        52.76 ns/op	      64 B/op	       1 allocs/op
/newlistbenchmarkCopy:_random_string:_len(8)-8         	21419151	        56.14 ns/op	     128 B/op	       1 allocs/op
/newlistbenchmarkSetItem:_random_string:_len(8)-8      	18315530	        66.10 ns/op	     128 B/op	       1 allocs/op
/newlistbenchmarkAppend:_random_string:_len(8)-8       	18125164	        66.52 ns/op	     128 B/op	       1 allocs/op
/newlistbenchmarkCopy:_random_string:_len(16)-8        	15583572	        77.38 ns/op	     256 B/op	       1 allocs/op
/newlistbenchmarkSetItem:_random_string:_len(16)-8     	12736520	        95.36 ns/op	     256 B/op	       1 allocs/op
/newlistbenchmarkAppend:_random_string:_len(16)-8      	12394372	        97.24 ns/op	     256 B/op	       1 allocs/op
/newlistbenchmarkCopy:_random_string:_len(32)-8        	10607676	       112.6 ns/op	     512 B/op	       1 allocs/op
/newlistbenchmarkSetItem:_random_string:_len(32)-8     	 7985698	       150.9 ns/op	     512 B/op	       1 allocs/op
/newlistbenchmarkAppend:_random_string:_len(32)-8      	 7889284	       151.4 ns/op	     512 B/op	       1 allocs/op
/newlistbenchmarkCopy:_random_string:_len(64)-8        	 6427102	       186.2 ns/op	    1024 B/op	       1 allocs/op
/newlistbenchmarkSetItem:_random_string:_len(64)-8     	 4379654	       271.7 ns/op	    1024 B/op	       1 allocs/op
/newlistbenchmarkAppend:_random_string:_len(64)-8      	 4494253	       268.9 ns/op	    1024 B/op	       1 allocs/op
/newlistbenchmarkCopy:_random_string:_len(128)-8       	 3435316	       348.6 ns/op	    2048 B/op	       1 allocs/op
/newlistbenchmarkSetItem:_random_string:_len(128)-8    	 2458584	       493.8 ns/op	    2048 B/op	       1 allocs/op
/newlistbenchmarkAppend:_random_string:_len(128)-8     	 2405611	       480.0 ns/op	    2048 B/op	       1 allocs/op
/newlistbenchmarkCopy:_random_string:_len(256)-8       	 1865565	       651.8 ns/op	    4096 B/op	       1 allocs/op
/newlistbenchmarkSetItem:_random_string:_len(256)-8    	 1000000	      1069 ns/op	    4096 B/op	       1 allocs/op
/newlistbenchmarkAppend:_random_string:_len(256)-8     	 1000000	      1060 ns/op	    4096 B/op	       1 allocs/op
*/
