package main

import (
	"fmt"

	"
)

var global interface{}

func main() {
	// const n = 1<<24 - 1
	const n = 1<<16 - 1

	k := make(kinds.KindMap, n)

	for i := 0; i < n; i++ {
		kn := kinds.RandomKind(true)
		k[kn.String()]++
	}

	s := kinds.GetEncodedString(n)

	fmt.Println("Encoded String: ", s[:20], " ... ", s[len(s)-20:])

	fmt.Println(k)

	fmt.Printf("map.Min():     %10.3d\n", k.Min())
	fmt.Printf("map.Max():     %10.3d\n", k.Max())
	fmt.Printf("map.Mean():    %10.3f\n", k.Mean())
	fmt.Printf("map.StDev():   %10.3f\n", k.StDev())

	// times := make([]testes.DataPoint, 0, 16)

	// for i := 0; i < n; i++ {
	// 	times = append(times, testes.GetData())
	// }

	// fmt.Println("Times: ", times)

}
