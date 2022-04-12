package greek

// import (
// 	"math/rand"
// 	"testing"
// )

// var (
// // tRun     = benchmark.TRun
// // tName    = benchmark.TName
// // tTypeRun = benchmark.TTypeRun
// )

// func Test_RandomGreekString(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		smoketest := RandomGreekString(i + 10)
// 		tRun(t, "randomGreekString", smoketest, smoketest)
// 	}
// }

// func Test_RandomGreek(t *testing.T) {
// 	for i := 0; i < 30; i++ {
// 		min := 10 + rand.Intn(20)       // 10 - 30
// 		max := min + 10 + rand.Intn(20) // 20 - 60 (> min+10)
// 		smoketest := RandomGreek(i+10, min, max)
// 		tRun(t, "randomGreek", smoketest, smoketest)
// 	}
// }

// func Test_charMap_toLower(t *testing.T) {
// 	type test struct {
// 		name string
// 		arg  string
// 		want rune
// 	}

// 	tests := []test{}

// 	for k, v := range Greek {
// 		tests = append(tests, test{k, k, v.lower})
// 	}

// 	tests = append(tests, test{"fake", "Zero", ReplacementChar})

// 	for _, tt := range tests {
// 		name := tName("Greek.toLower", "Greek letter for "+tt.name, "")
// 		tTypeRun(t, name, Greek.toLower(tt.arg), tt.want)
// 	}
// }
// func Test_charMap_toUpper(t *testing.T) {
// 	type test struct {
// 		name string
// 		arg  string
// 		want rune
// 	}

// 	tests := []test{}

// 	for k, v := range Greek {
// 		tests = append(tests, test{k, k, v.upper})
// 	}

// 	tests = append(tests, test{"fake", "Zero", ReplacementChar})

// 	for _, tt := range tests {
// 		name := tName("Greek.toUpper", "Greek letter for "+tt.name, "")
// 		tTypeRun(t, name, Greek.toUpper(tt.arg), tt.want)
// 	}
// }

// func Test_charMap_ToLower(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		s := RandomGreekString(i + 10)
// 		smoketest := Greek.ToLower(s)
// 		tRun(t, "Greek.ToLower", smoketest, smoketest)
// 	}
// }

// func Test_charMap_ToUpper(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		s := RandomGreekString(i + 10)
// 		smoketest := Greek.ToUpper(s)
// 		tRun(t, "Greek.ToLower", smoketest, smoketest)
// 	}
// }
