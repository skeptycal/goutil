package main

import (
	"fmt"
	"unicode"

	"
)

type Booler = defaults.Booler

var (
	pf = fmt.Printf
	ab = defaults.AnyBooler

	samples = []struct {
		name string
		b    Booler
	}{
		{"boolean true", ab(true)},
		{"boolean false", ab(false)},
		{"integer 0", ab(int(0))},
		{"integer 42", ab(int(42))},
		{"empty string", ab("")},
		{"'false' string", ab("false")},
		{"'true' string", ab("true")},
		{"'0' string", ab("0")},
		{"nil", ab(nil)},
		{"empty slice", ab([]rune{})},
		{"non-empty slice", ab([]rune{'4', '2'})},
		{"unicode.IsUpper('G')", ab(unicode.IsUpper('G'))},
		{"unicode.IsUpper('g')", ab(unicode.IsUpper('g'))},
	}
)

func main() {

	fmt.Println("False Example Values:")
	for _, ss := range samples {
		if !ss.b.AsBool() {
			pf("%-40s : %v\n", "AnyBool("+ss.name+")", ss.b)
		}
	}

	fmt.Println("")
	fmt.Println("True Example Values:")

	for _, ss := range samples {
		if ss.b.AsBool() {
			pf("%-40s : %v\n", "AnyBool("+ss.name+")", ss.b)
		}
	}
}
