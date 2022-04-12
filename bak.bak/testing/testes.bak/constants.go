package testes

// UnicodeReplacementChar is the recognized unicode
// replacement character for malformed unicode or
// errors in encoding.
//
// It is also found in unicode.UnicodeReplacementChar
const UnicodeReplacementChar rune = '\uFFFD'

// Encoding, searching, and parsing constants.
const (
	AllUppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AllLowercase    = "abcdefghijklmnopqrstuvwxyz"
	AllDigits       = "0123456789"
	AllAlpha        = AllLowercase + AllUppercase
	AllAlphanumeric = AllAlpha + AllDigits

	// Base64 encoding (rfc4648) - padding is done at the end with '='
	// and this character should not appear anywhere else.
	//
	// Reference: https://datatracker.ietf.org/doc/html/rfc4648
	Base64Set = AllUppercase + AllLowercase + AllDigits + "+/=" // rfc4648

	// rfc4648 - padding is done at the end with '='
	// and this character should not appear anywhere else.
	//
	// Reference: https://datatracker.ietf.org/doc/html/rfc4648
	Base64urlSet = AllUppercase + AllLowercase + AllDigits + "-_=" // rfc4648
)

// ASCII format
// Reference: https://datatracker.ietf.org/doc/html/rfc20
