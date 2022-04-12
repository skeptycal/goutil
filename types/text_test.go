package types_test

// import (
// 	"testing"

// 	"
// )

// func Test_byteMap_Len(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		b    *byteMap
// 		want int
// 	}{
// 		{"abcde", NewByteMap("abcde"), 5},
// 		{"0123456789", NewByteMap("0123456789"), 10},
// 		{"AAAAA", NewByteMap("AAAAA"), 1},
// 		{"AABAA", NewByteMap("AABAA"), 2},
// 	}
// 	for _, tt := range tests {
// 		tRun(t, tt.name, tt.b.Len(), tt.want)
// 	}
// }

// var textTests = []benchmark.Tester{
// 	{"abcde", "abcde", len(Frequency("abcde")), 5, false},
// 	{"0123456789", "0123456789", len(Frequency("0123456789")), 10, false},
// 	{"AAAAA", "AAAAA", len(Frequency("AAAAA")), 1, false},
// 	{"AABAA", "AABAA", len(Frequency("AABAA")), 2, false},
// 	{"BBB", "BBB", len(Frequency("BBB")), 3, true}, // incorrect - want error
// }

// func TestFrequency(t *testing.T) {
// 	benchmark.NewTestSet(t, "TestFrequency", textTests).Run()
// }
