package main

import (
	"fmt"

	generic "github.com/skeptycal/goutil/generic/generictypes/dict"
)

func main() {
	d := generic.Dict[string, int]{}

	fmt.Println("Is Empty? ", d.IsEmpty())
	fmt.Println(" /puts stuff in the map ......")
	d.Set("one", 1)
	d.Set("two", 4)
	d["three"] = 3
	fmt.Println(d)
	fmt.Println("Is Empty? ", d.IsEmpty())
	one, ok := d.Get("one")
	if !ok {
		fmt.Println("element 1 not found")
	} else {
		fmt.Println("d[\"one\"]: ", one)
	}
	fmt.Println("length of map: ", d.Len())
	fmt.Println(" /fixes the value of 'two' ......")
	d.Set("two", 2)
	fmt.Println(d)

}
