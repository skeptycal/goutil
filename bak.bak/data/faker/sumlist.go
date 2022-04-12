package faker

import "fmt"

type sumList struct {
	n    int
	list []float64
}

func (s sumList) Get(i int) float64    { return s.list[i] }
func (s sumList) Set(i int, v float64) { s.list[i] = v }
func (s sumList) Append(v float64)     { s.list = append(s.list, v) }
func (s sumList) Len() int             { return len(s.list) }
func (s sumList) Swap(i, j int)        { s.list[i], s.list[j] = s.list[j], s.list[i] }
func (s sumList) Less(i, j int) bool   { return s.list[i] < s.list[j] }
func (s sumList) Reset()               { s.list = s.list[:0] }
func (s sumList) String() string       { return fmt.Sprintf("%v runs of %v : %v\n", s.Len(), s.n, s.list) }
