package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	pathFlag := "."
	sizeFlag := true
	outputSEP := "\n"

	if len(os.Args) > 1 {
		pathFlag = os.Args[1]
	}

	// fmt.Printf("listing for %s\n", pathFlag)
	dirs, err := GetDirs(pathFlag, true, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("dir list:")
	for _, dir := range dirs {
		fmt.Printf(" %v", dir.Name())
	}
	fmt.Println()

	for _, fi := range dirs {
		s := ""
		if fi.IsDir() {
			s = fi.Name()

			if sizeFlag {
				s += fmt.Sprintf(" %d", fi.Size())
			}

			s += outputSEP

			// fmt.Printf("%v %s %d\n", fi.Mode(), fi.Name(), fi.Size())
		}
		fmt.Print(s)

	}
	fmt.Println()

	for i := 0; i < 5; i++ {
		fmt.Println(ansiFmt(i))
	}

}
