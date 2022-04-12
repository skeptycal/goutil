package main

import (
	"fmt"
	"
	"log"
	// "github.com/integralist/go-findroot/find"
)

func main() {
	root, err := find.Repo()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	// fmt.Printf("%+v", root)
	fmt.Printf("%v", root)
	// {Name:go-findroot Path:/Users/M/Projects/golang/src/github.com/integralist/go-findroot}
}
