package parser

type (
	// Any type of data
	Any interface{}
	// StringFunction takes a string and returns a string
	StringFunction = func(s string) string
)

// GetSetter implements basic generic get and set methods.
// May be used to wrap a dictionary, INI config file,
// TOML file, JSON config file, struct, etc ...
type GetSetter interface {
	Get(key Any) (value Any, err error)
	Set(key, value Any) (err error)
}

// Option performs a single operation on text data and
// returns the result.
type Option interface {
	Parse(string) string
}

// Parameter stores a single parameter and returns the
// name, value, and functionality as needed. Any type
// may be stored as a parameter, including functions
// that provide custom functionality.
//
// e.g. a parameter named "newline" might return the
// string that is used as a linebreak in the text.
//  nl.Name()
//  // "newline"
//  nl.Function()
//  func()  { return "\n" }
//
// a parameter named "upper" might return the function
// that is used to change text to upper case.
//  upper.Name()
//  // "UpperCase"
//  upper.Function()
//  func(s string) string { return strings.ToUpper(s) }
type Parameter interface {
	Name() string
	Function(name string) Any
}

type ParseOptioner interface {
	GetSetter
}

// ParserOptions contains the configuration information
// for the Parser.
//
// These fields are only used if the behavior is desired,
// otherwise they should not be set.
//
// If lineseps are provided, they will be converted to the
// default linesep (or nl if it is given)
type ParserOptions struct {
	nl            string            // single newline character to use (default \n)
	sep           string            // single field delimiter to use (default \t)
	characterCase Cases             // none, upper, lower, title, camel, snake, Pascal, kehab
	lineseps      []string          // list of acceptable line delimiters (default empty)
	fieldseps     []string          // list of acceptable field delimiters (default empty)
	prefix        string            // prefix to remove (default "")
	addprefix     string            // prefix to add (default "")
	suffix        string            // suffix to remove (default "")
	addsuffix     string            // suffix to add (default "")
	cut           []string          // strings to remove (default empty)
	replace       map[string]string // map of strings to replace (default empty)
}
