package main

import (
	"fmt"
	"log"
	"os"
)

const pattern = "."

func main() {

	// As of Go 1.16, os.ReadDir is a more efficient and correct choice: it returns a list of fs.DirEntry instead of fs.FileInfo, and it returns partial results in the case of an error midway through reading a directory.
	files, err := os.ReadDir(pattern)
	// files, err = ioutil.ReadDir(pattern)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(files)

	// var fs fileSize
	// for _, f := range files {
	// 	fi, _ := f.Info()
	// 	fs.int64 = fi.Size()
	// 	fmt.Printf("%s %6v %s\n", f.Type(), fs, f.Name())

	// }
}
