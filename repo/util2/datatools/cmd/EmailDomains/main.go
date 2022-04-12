package main

import (
	"fmt"
	"os"
	"strings"

	"
)

func main() {
    list := strings.Join(os.Args[1:]," ")
    out := format.GetDomainNames(list)
    fmt.Println(out)
}
