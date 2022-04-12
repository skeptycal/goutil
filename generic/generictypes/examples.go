package generic

import "fmt"

func ExampleAddOne[T Number](n T) {
	i := n + 1
	fmt.Println(i)
}
