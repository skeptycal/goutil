package main

import (
	"fmt"
	"os"
)

const skipCharacter = "#"

type Processer interface {
	Process() error
	String() string
}

func In(needles, haystack string) bool {
	for _, needle := range needles {
		for _, straw := range haystack {
			if needle == straw {
				return true
			}
		}
	}
	return false
}

func main() {
	filename := os.Args[1]

	f, err := memfile.NewMemFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating memfile: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(f)

	var b []byte

	_, err = f.Read(b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading memfile: %v\n", err)
		os.Exit(1)
	}

	p, err := memfile.NewParser(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error initializing parser: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(p)

	// for _, line := range strings.Split(string(b), "\n") {
	// 	if strings.HasPrefix(line, skipCharacter) {
	// 		continue
	// 	}
	// 	// process line
	// }
}
