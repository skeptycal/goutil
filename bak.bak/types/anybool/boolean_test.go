package anybool

import (
	"testing"
	"unicode"
)

/* Benchmark results for Booler (bool value vs interface value)

Takeaways:
- bools are very, very fast
- pointers are very fast
- ints, uints, floats, struct{}, maps, slices(not []byte), channels are fast
- strings and []byte are half speed of above (even other slices???)
- arrays are noticeably slower than all others

BenchmarkBooler/NewBool(true)-8         	196670611	         5.918 ns/op	   0 B/op	       0 allocs/op
BenchmarkBooler/AnyBool(bool)-8         	 34330753	        47.42  ns/op	   0 B/op	       0 allocs/op

BenchmarkBooler/AnyBool(false)-8           165912350	         7.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/AnyBool(true)-8            163697052	         7.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/AnyBool(true)-8          	55380460	        22.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/the_empty_string-8       	26632684	        51.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/'false'-8                	19048916	        58.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/'0'-8                    	26050334	        59.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/true_string-8            	24274688	        55.49 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/empty_[]byte-8           	22117113	        61.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/non-empty_[]byte-8       	22498663	        62.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/int_0-8                  	30002030	        48.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/int8_0-8                 	28536974	        43.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/int16_0-8                	32840272	        38.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/int32_0-8                	42892950	        29.70 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/int64_0-8                	40022901	        29.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/int_42-8                 	37534357	        31.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/int8_42-8                	42608340	        30.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/int16_42-8               	40746220	        28.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/int32_42-8               	42150926	        28.81 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/int64_42-8               	37750736	        28.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/uint_0-8                 	49273642	        24.25 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/uint8_0-8                	51577038	        29.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/uint16_0-8               	41422041	        26.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/uint32_0-8               	48347780	        24.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/uint64_0-8               	46554181	        27.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/uint_42-8                	42202501	        27.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/uint8_42-8               	46550343	        26.96 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/uint16_42-8              	38536478	        31.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/uint32_42-8              	39620086	        30.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/uint64_42-8              	34818856	        37.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/float_0.0-8              	41392989	        26.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/float32_0.0-8            	41017986	        29.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/float64_0.0-8            	41653528	        25.66 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/float_42.0-8             	44531652	        27.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/float32_42.0-8           	46420838	        27.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/float64_42.0-8           	30215378	        33.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/empty_struct_(thing{})-8 	32121490	        36.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/false_struct-8           	24694210	        45.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/true_struct-8            	25428420	        48.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/empty_slice-8            	35387748	        35.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/slice-8                  	33835500	        35.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/empty_array-8            	11728260	       102.3  ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/array-8                  	14829880	        70.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/empty_map-8              	26005053	        41.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/map[int]byte-8           	32912179	        36.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/map[string]string-8      	33690870	        37.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/empty_chan-8             	32297919	        36.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/nil_pointer-8            	77658804	        15.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/true_pointer_(not_nil)-8 	76069932	        15.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkBooler/false_pointer_(not_nil)-8   81939463	        15.67 ns/op	       0 B/op	       0 allocs/op
*/

/* Benchmark results for AsBool(), the function that returns the boolean value
and likely the most common use case. e.g.

 if AnyBooler(struct{some_stuff}) {
	// do stuff if true ...
 }

AnyBool(true)-8		   		  459922761		   		      2.588 ns/op
AnyBool(false)-8		   	  403378915		   		      3.181 ns/op
nil-8		   		      	  341665299		   		      4.176 ns/op
nil_pointer-8		   		  194859840		   		      6.159 ns/op
false_pointer_(not_nil)-8	  194509022		   		      6.159 ns/op
true_pointer_(not_nil)-8	  192762529		   		      6.167 ns/op
uint8_0-8		   		      121829476		   		      9.819 ns/op
int8_0-8		   		      122175487		   		      9.833 ns/op
uint_0-8		   		      121726291		   		      9.835 ns/op
float64_0.0-8		   		  121863541		   		      9.853 ns/op
float_0.0-8		   		      121720087		   		      9.860 ns/op
float32_0.0-8		   		  121386322		   		      9.888 ns/op
int_0-8		   		      	  100000000		   		      10.00 ns/op
int64_0-8		   		      100000000		   		      10.14 ns/op
int32_0-8		   		      100000000		   		      10.21 ns/op
uint64_0-8		   		      121869430		   		      10.57 ns/op
uint16_0-8		   		      121591849		   		      11.11 ns/op
uint32_0-8		   		      121507041		   		      11.26 ns/op
int16_0-8		   		      100000000		   		      11.43 ns/op
int8_42-8		   		      100000000		   		      11.54 ns/op
int16_42-8		   		      100000000		   		      11.55 ns/op
int_42-8		   		      100000000		   		      11.59 ns/op
int32_42-8		   		      100000000		   		      11.93 ns/op
uint_42-8		   		      100000000		   		      11.96 ns/op
uint32_42-8		   		       99113143		   		      12.00 ns/op
uint64_42-8		   		       99979520		   		      12.04 ns/op
float32_42.0-8		   	       99430346		   		      12.05 ns/op
uint16_42-8		   		       97423687		   		      12.07 ns/op
uint8_42-8		   		      100000000		   		      12.12 ns/op
empty_struct_(thing{})-8	   99363106		   		      12.15 ns/op
int64_42-8		   		       97503504		   		      12.27 ns/op
float_42.0-8		   		   96379418		   		      12.44 ns/op
float64_42.0-8		   		   96647210		   		      12.62 ns/op
the_empty_string-8		   	   95115426		   		      12.66 ns/op
true_string-8		   		   82432443		   		      14.55 ns/op
'0'-8		   		      	   81981682		   		      14.65 ns/op
'false'-8		   		       81403990		   		      14.72 ns/op
empty_[]byte-8		   		   73341975		   		      16.39 ns/op
non-empty_[]byte-8		   	   72976795		   		      16.40 ns/op
slice-8		   		      	   73297732		   		      16.40 ns/op
empty_slice-8		   		   73147689		   		      16.43 ns/op
empty_map-8		   		       68188597		   		      17.60 ns/op
map[string]string-8		   	   68090418		   		      17.63 ns/op
map[int]byte-8		   		   67848358		   		      17.71 ns/op
empty_chan-8		   		   67375986		   		      17.84 ns/op
false_struct-8		   		   48710113		   		      24.76 ns/op
array-8		   		      	   47698861		   		      25.29 ns/op
true_struct-8		   		   45950602		   		      26.55 ns/op
empty_array-8		   		   22106232		   		      54.34 ns/op
*/

var true_bool bool = true
var truePtr *bool = &true_bool

var false_bool bool = false
var falsePtr *bool = &false_bool
var nilPtr *uintptr = nil
var trueBooler = AnyBooler(true)

var retval Any = ""

var boolerTests = []struct {
	name  string
	input Booler
	want  string
}{
	// nil
	{"nil", AnyBooler(nil), "false"},

	// Booler
	// TODO - this test is not working ...
	// {"Booler", AnyBooler(AnyBooler(true)), "true"},
	{"Booler", AnyBooler(AnyBooler(false)), "false"},

	// bools
	{"false", NewBooler(false), "false"},
	{"true", NewBooler(true), "true"},
	{"AnyBool(false)", AnyBooler(false), "false"},
	{"AnyBool(true)", AnyBooler(true), "true"},

	// strings
	{"the empty string", AnyBooler(""), "false"},
	{"'false'", AnyBooler("false"), "false"},
	{"'0'", AnyBooler("0"), "false"},
	{"true string", AnyBooler("true"), "true"},
	{"fake string", AnyBooler("fake"), "false"},

	// []byte
	{"empty []byte", AnyBooler([]byte("")), "false"},
	{"non-empty []byte", AnyBooler([]byte("fake")), "true"},

	// ints
	{"int 0", AnyBooler(int(0)), "false"},
	{"int8 0", AnyBooler(int8(0)), "false"},
	{"int16 0", AnyBooler(int16(0)), "false"},
	{"int32 0", AnyBooler(int32(0)), "false"},
	{"int64 0", AnyBooler(int64(0)), "false"},
	{"int 42", AnyBooler(int(42)), "true"},
	{"int8 42", AnyBooler(int8(42)), "true"},
	{"int16 42", AnyBooler(int16(42)), "true"},
	{"int32 42", AnyBooler(int32(42)), "true"},
	{"int64 42", AnyBooler(int64(42)), "true"},

	// uints
	{"uint 0", AnyBooler(uint(0)), "false"},
	{"uint8 0", AnyBooler(uint8(0)), "false"},
	{"uint16 0", AnyBooler(uint16(0)), "false"},
	{"uint32 0", AnyBooler(uint32(0)), "false"},
	{"uint64 0", AnyBooler(uint64(0)), "false"},
	{"uint 42", AnyBooler(uint(42)), "true"},
	{"uint8 42", AnyBooler(uint8(42)), "true"},
	{"uint16 42", AnyBooler(uint16(42)), "true"},
	{"uint32 42", AnyBooler(uint32(42)), "true"},
	{"uint64 42", AnyBooler(uint64(42)), "true"},

	// floats
	{"float 0.0", AnyBooler(0.0), "false"},
	{"float32 0.0", AnyBooler(float32(0.0)), "false"},
	{"float64 0.0", AnyBooler(float64(0.0)), "false"},
	{"float 42.0", AnyBooler(42.0), "true"},
	{"float32 42.0", AnyBooler(float32(42.0)), "true"},
	{"float64 42.0", AnyBooler(float64(42.0)), "true"},

	// complex
	{"complex (0,0)", AnyBooler(complex(0, 0)), "false"},
	{"complex (42,1)", AnyBooler(complex(42, 1)), "true"},

	// structs
	{"empty struct (thing{})", AnyBooler(Cosa{}), "false"},
	{"false struct", AnyBooler(struct{ i int }{i: 0}), "false"},
	{"true struct", AnyBooler(struct{ i int }{i: 1}), "true"},

	// slices
	{"empty slice", AnyBooler([]rune{}), "false"},
	{"slice", AnyBooler([]rune{'4', '2'}), "true"},

	// arrays
	{"empty array", AnyBooler([4]byte{}), "false"},
	{"array", AnyBooler([2]byte{'4', '2'}), "true"},

	// maps
	{"empty map", AnyBooler(make(map[string]string)), "false"},
	{"map[int]byte", AnyBooler(map[int]byte{1: '4', 2: '2'}), "true"},
	{"map[string]string", AnyBooler(map[string]string{"4": "2"}), "true"},

	// channels
	{"empty chan", AnyBooler(make(chan string)), "false"},

	// pointers
	{"nil pointer", AnyBooler(nilPtr), "false"},
	{"true pointer (not nil)", AnyBooler(truePtr), "true"},
	{"false pointer (not nil)", AnyBooler(falsePtr), "true"},

	// funcs
	{"unicode.IsUpper('J')", AnyBooler(unicode.IsUpper('J')), "true"},
	{"unicode.IsUpper('J')", AnyBooler(unicode.IsUpper('j')), "false"},
}

var newBoolTests = boolerTests[1:2]

func BenchmarkBooler(b *testing.B) {
	for _, bb := range boolerTests {
		b.Run(bb.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				retval = bb.input.String()
			}
		})
	}
}

func BenchmarkBoolerAsBool(b *testing.B) {
	for _, bb := range boolerTests {
		b.Run(bb.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				retval = bb.input.AsBool()
			}
		})
	}
}

func Test_boolean_Enable(t *testing.T) {
	want := "true"
	for _, tt := range newBoolTests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.input
			b.Enable()
			if got := b.String(); got != want {
				t.Errorf("AnyBool(%v) = %v, want %v", tt.name, got, want)
			}
		})
	}
}

func Test_boolean_Disable(t *testing.T) {
	want := "false"
	for _, tt := range newBoolTests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.input
			b.Disable()
			if got := b.String(); got != want {
				t.Errorf("NewBool(%v) = %v, want %v", tt.name, got, want)
			}
		})
	}
}

func TestAnyBool(t *testing.T) {
	for _, tt := range boolerTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.String(); got != tt.want {
				t.Errorf("AnyBool(%v) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func Test_anyBool_Enable(t *testing.T) {
	want := "true"

	for _, tt := range boolerTests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.input
			b.Enable()
			if got := b.String(); got != want {
				t.Errorf("NewBool(%v) = %v, want %v", tt.name, got, want)
			}
		})
	}
}

func Test_anyBool_Disable(t *testing.T) {
	want := "false"

	for _, tt := range boolerTests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.input
			b.Disable()
			if got := b.String(); got != want {
				t.Errorf("NewBool(%v) = %v, want %v", tt.name, got, want)
			}
		})
	}
}

func Test_boolean_AsBool(t *testing.T) {
	type fields struct {
		bool bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &boolean{
				bool: tt.fields.bool,
			}
			if got := b.AsBool(); got != tt.want {
				t.Errorf("boolean.AsBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTrue(t *testing.T) {
	tests := []struct {
		name string
		v    Any
		want bool
	}{
		{"0", 0, false},
		{"42", 42, true},
		{`"true"`, true, true},
		{`"false"`, false, false},
		{`[]byte{}`, []byte{}, false},
		{`[]byte("42")`, []byte("42"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTrue(tt.v); got != tt.want {
				t.Errorf("IsTrue(%v) = %v, want %v", tt.name, got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFalse(tt.v); got != !tt.want {
				t.Errorf("IsFalse(%v) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestNewBooler(t *testing.T) {
	want := true
	got := NewBooler(true)
	if got.AsBool() != want {
		t.Errorf("NewBooler(true) = %v, want %v", got, want)
	}

	want = false
	got.Disable()
	if got.AsBool() != want {
		t.Errorf("NewBooler(true) = %v, want %v", got, want)
	}

	got = NewBooler(false)
	if got.AsBool() != want {
		t.Errorf("NewBooler(true) = %v, want %v", got, want)
	}

	want = true
	got.Enable()
	if got.AsBool() != want {
		t.Errorf("NewBooler(true) = %v, want %v", got, want)
	}
}
