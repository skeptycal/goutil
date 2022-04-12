package sigfig

import (
	"fmt"
	"math/big"
)

type (
	sfFloat big.Float

	Floater interface {
		Sigfigs() int
		Value() float64
		Unit() string
		fmt.Stringer
	}
)

func sfMultiply(values ...Floater) (retval int) {
	// calculate sigfigs for multiplication / division

	return 0
}

func sfAdd(value ...Floater) (retval int) {
	// calculate sigfigs for addition / subtraction

	return 0
}

func sigfig(value float64) int {
	// calculate general sigfigs

	return 0
}

func (m *measurement) Sigfigs() int {
	if m.sigfigs == 0 {
		m.sigfigs = sigfig(m.value.Value())
	}
	return m.sigfigs
}

func (m *measurement) Add(other ...measurement) (retval measurement) {
	for _, o := range other {
		m.Value
	}
	self := m.value

}
