package sigfig

import (
	"math/big"
	mrand "math/rand"
)

// ... "The zero (uninitialized) value for a Float is ready to use
// and represents the number +0.0 exactly, with precision 0
// and rounding mode ToNearestEven."
//
// reference: Go 1.17.5 math/big/float.go
var zero *decimal = &decimal{}

type (
	decimal = big.Float

	// A decimal represents an unsigned floating-point number in decimal representation.
	// The value of a non-zero decimal d is d.mant * 10**d.exp with 0.1 <= d.mant < 1,
	// with the most-significant mantissa digit at index 0. For the zero decimal, the
	// mantissa length and exponent are 0.
	// The zero value for decimal represents a ready-to-use 0.0.
	//
	// (from big/decimal in Go 1.17.5 SL)
	// decimal struct {
	// 	mant []byte // mantissa ASCII digits, big-endian
	// 	exp  int    // exponent
	// 	buf  []byte // current output based on 'sf' digits (buffered JIT)
	// 	sf   int    // current number of sigfigs desired
	// }

	measurement struct {
		value   Floater
		sigfigs int
		unit    string
	}
)

func (d *decimal) parse(s string) *decimal {

}

func NewDec(s string) *decimal {
	d := &decimal{}.parse(s)
	d.parse(s)
}

func r(n int) int { return mrand.Intn(n) }

// rand returns a random decimal with a max length of n.
func rand(n int) *decimal {

	if n == 0 {
		return zero
	}
	pctDP := 80

	var s string

	d := new(decimal)

	if n < 2 {

	}
	d.Parse(s)

	// add decimal point if random chance roll hits and length is > 2
	dp := r(100) < pctDP && n > 2

	n1 := r()

	if dp {
		var dploc int = n

		dploc = mrand.Intn(n-2) + 1
	}

}
