// Package format contains functions that format numeric values.
package format

import (
	"strings"
)

type stringWriter struct {
	strings.Builder
	intpart  string
	decpart  string
	exponent string
}

func (w *stringWriter) space() { w.WriteString(" ") }
func (w *stringWriter) dot()   { w.WriteString(".") }
func (w *stringWriter) exp()   { w.WriteString(" e" + w.exponent) }
func (w *stringWriter) parse(value string) {
	value = strings.TrimSpace(value)
	value = strings.ToLower(value)
	mantissa := value

	eloc := strings.Index(value, "e")
	if eloc != -1 {
		mantissa = value[:eloc]
		w.exponent = value[eloc+1:]
	}

	w.intpart = mantissa
	w.decpart = ""

	dloc := strings.Index(mantissa, ".")
	if dloc > 0 {
		w.intpart = mantissa[:dloc]
		w.decpart = mantissa[dloc+1:]
	}

	w.loadString()
}
func (w *stringWriter) loadString() {

	if w.intpart[0] == '=' {
		w.WriteByte(w.intpart[0])
		w.intpart = w.intpart[1:]
	}

	rem := len(w.intpart)%3 + 2

	t := ""
	j := len(w.intpart) - 1
	for i := j; i >= 0; i-- {
		t += string(w.intpart[i])
		// w.WriteByte(w.intpart[i])
		if (i+rem)%3 == 0 && i < len(w.intpart) {
			t += " "
			// w.space()
		}
	}

	// w.WriteByte(w.intpart[len(w.intpart)-1])

	if w.decpart != "" {
		w.dot()

		for i := 0; i < len(w.decpart); i++ {
			w.WriteByte(w.decpart[i])
			if (i+1)%3 == 0 {
				w.space()
			}

		}
	}
	if w.exponent != "" {
		w.exp()
	}
}

// Reverse is not utf8 compatible
// ... focus on
//  utf8.DecodeRuneInString(s string) (r rune, size int)
func Reverse(s string) string {
    sb := strings.Builder{}
    defer sb.Reset()

	for i := len(s) - 1; i > -1; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}

// Reverse2 reverses
// Reference: https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func Reverse2(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Reverse3 reverses
// Reference: https://stackoverflow.com/a/1754209
func Reverse3(s string) string {
	// Get Unicode code points.
	n := 0
	runes := make([]rune, len(s))
	for _, r := range s {
		runes[n] = r
		n++
	}
	runes = runes[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	// Convert back to UTF-8.
	return string(runes)
}

// Reverse4 reverses
// combines the best from SO answers (Reverse2 and Reverse3)
func Reverse4(s string) string {
	runes := []rune(s)
	n := len(runes)
	// Reverse
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

// Reverse5 returns a string with the bytes of s in reverse order.
func Reverse5(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for i := len(s) - 1; i >= 0; i-- {
		b.WriteByte(s[i])
	}
	return b.String()
}

// ReverseRune returns a string with the runes of s in reverse order.
// Invalid UTF-8 sequences, if any, will be reversed byte by byte.
func ReverseRune(s string) string {
	res := make([]byte, len(s))
	prevPos, resPos := 0, len(s)
	for pos := range s {
		resPos -= pos - prevPos
		copy(res[resPos:], s[prevPos:pos])
		prevPos = pos
	}
	copy(res[0:], s[prevPos:])
	return string(res)
}

// func Reverse5(s string) string {
// 	rb := runebuilder.Builder{}
// 	for i := len(s) - 1; i > -1; i-- {
// 		rb.WriteRune(r)
// 	}
// 	return rb.String()
// }

// func Reverse6(s string) string {
// 	rb := runebuilder.Builder{}
// 	size := utf8.RuneCountInString(s)
// 	for i := size - 1; i >= 0; i-- {
// 		utf8.DecodeRuneInString(s)
// 		rb.WriteRune()
// 	}
// 	return rb.String()
// }

// func Reverse8(s string) string {

// 	rb := runebuilder.Builder{}
// 	rb.ReSize(len(s))
// 	rb.Grow(len(s))
// 	fmt.Println("len: ", rb.Len())

// 	size := len(s) - 1
// 	for i, r := range s {
// 		// fmt.Printf("i: %d    r: %v", i, r)
// 		rb.Put(r, size-i)
// 	}
// 	return rb.String()
// }

// NumSpace formats numeric values for readability by adding
// spaces every three digits.
//
// e.g.
//
//    12345678.87654321 e-42
// first, split off any exponent from the mantissa
//
//    12345678.87654321  and  e-42
// next, split off any decimal part from the integer part
//    12345678   and    .87654321
//
// next, add spaces between digits in the integer part
//
//    12 345 678  and    .876 543 21
//
// finally, add any exponent back to the mantissa
func NumSpace(s string) string {

	sb := &stringWriter{strings.Builder{}, "", "", ""}
	sb.parse(s)

	return sb.String()

}

// python version
/* # Take a very large number and pretty print it in triplets of 3 digits, each triplet separated by a space.
def pnum_spc(n): print(' '.join([''.join(list(str(n))[::-1][i:i+3]) for i in range(0, len(str(n)), 3)][::-1]))
# >>> pnum_spc(32 ** 13)
# 36 893 488 147 419 103 232
*/

// python version
/* # Print numbers as 32-bit binary numbers w/ spaces giving 4-bit words
def pbin_spc(n): print(' '.join([''.join(list(f'{n:032b}')[::-1][i:i+4][::-1]) for i in range(0, len(f'{n:032b}'), 4)][::-1]))
# >>> "{0:032b}".format(1234)
# '00000000000000000000010011010010'
# >>> pbin_spc(1234)
# 0000 0000 0000 0000 0000 0100 1101 0010
*/
