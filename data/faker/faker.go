package faker

const (

	// ReplacementChar is the recognized unicode replacement
	// character for malformed unicode or errors in
	// encoding.
	//
	// It is also found in unicode.ReplacementChar
	ReplacementChar rune = '\uFFFD'

	UPPER    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LOWER    = "abcdefghijklmnopqrstuvwxyz"
	DIGITS   = "0123456789"
	ALPHA    = LOWER + UPPER
	ALPHANUM = ALPHA + DIGITS
)

func NoOp(any interface{}) []AnyValue { return nil } // noop function
