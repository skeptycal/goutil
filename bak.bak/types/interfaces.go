package types

type (

	// Errer processes and returns an error
	Errer interface {
		Err() error
	}

	// Closer is the interface that wraps the basic Close method.
	//
	// The behavior of Close after the first call is undefined.
	// Specific implementations may document their own behavior.
	Closer interface {
		Close() error
	}

	// An Enabler represents an object that can be enabled or disabled.
	Enabler interface {
		Enable()
		Disable()
	}

	// Protector is used when an object needs to be protected or unprotected
	// (inspired by the "write-protect" tabs of floppy disks)
	Protector interface {
		Protect()
		Unprotect()
	}

	// A GetSetter represents an object that can be accessed using
	// Get and Set methods to access an underlying map.
	//
	// Example methods:
	//  func (p *padding) Get(key Any) (Any, error) {
	//		// TODO get the value that matches the key
	//  	return nil, nil
	//  }
	//  func (p *padding) Set(key, value Any) error {
	//		// TODO set the value that matches the key
	//  	return nil
	//  }
	GetSetter interface {
		Get(key Any) (Any, error)
		Set(key Any, value Any) error
	}

	// Slicer returns the slice of keys and values that are
	// asoociated with the underlying data structure.
	//
	// Example methods:
	// 	func (d *dict) Keys() []Any {
	//		// TODO return a list of keys
	// 		keys := make([]Any, len(d.m))
	//	 	for k := range d.m {
	//	 		keys = append(keys, k)
	//	 	}
	//	 	return keys
	//	 }
	//	 func (d *dict) Values() []Any {
	//		// TODO return a list of values
	//	 	values := make([]Any, len(d.m))
	//	 	for _, v := range d.m {
	//	 		values = append(values, v)
	//	 	}
	//	 	return values
	//	 }
	Slicer interface {
		Keys() []Any
		Values() []Any
	}

	// Printer implements common printing functions similar
	// to the standard library fmt package.
	//
	// Example methods:
	// 	func (p *padding) Print(args ...interface{}) (n int, err error) {
	//		// TODO Print unformatted args
	// 		return n, err
	// 	}
	// 	func (p *padding) Println(args ...interface{}) (n int, err error) {
	//		// TODO Print unformatted args with line break (NL)
	// 		return n, err
	// 	}
	// 	func (p *padding) Printf(format string, args ...interface{}) (n int, err error) {
	//		// TODO Print formatted args
	// 		return n, err
	// 	}
	Printer interface {
		Print(args ...interface{}) (n int, err error)
		Println(args ...interface{}) (n int, err error)
		Printf(format string, args ...interface{}) (n int, err error)
	}

	// Fprinter implements common printing functions similar
	// to the standard library fmt package.
	//
	// Example methods:
	// 	func (p *padding) Fprint(w Writer, args ...interface{}) (n int, err error) {
	//		// TODO Print unformatted args to Writer
	// 		return n, err
	// 	}
	// 	func (p *padding) Fprintln(w Writer, args ...interface{}) (n int, err error) {
	//		// TODO Print unformatted args to Writer with line break (NL)
	// 		return n, err
	// 	}
	// 	func (p *padding) Fprintf(w Writer, format string, args ...interface{}) (n int, err error) {
	//		// TODO Print formatted args to Writer
	// 		return n, err
	// 	}
	Fprinter interface {
		Fprint(w Writer, a ...interface{}) (n int, err error)
		Fprintln(w Writer, a ...interface{}) (n int, err error)
		Fprintf(w Writer, format string, a ...interface{}) (n int, err error)
	}

	// Sprinter implements common printing functions similar
	// to the standard library fmt package.
	//
	// Example methods:
	// 	func (p *padding) Sprint(args ...interface{}) string {
	//		// TODO Print unformatted args to string
	// 		return ""
	// 	}
	// 	func (p *padding) Sprintln(args ...interface{}) string {
	//		// TODO Print unformatted args to string with line break (NL)
	// 		return ""
	// 	}
	// 	func (p *padding) Sprintf(format string, args ...interface{}) string {
	//		// TODO Print formatted args to string
	// 		return ""
	// 	}
	Sprinter interface {
		Sprint(args ...interface{}) string
		Sprintln(args ...interface{}) string
		Sprintf(format string, args ...interface{}) string
	}
)

//* Interfaces from standard library for reference.
// Copied here to avoid larger dependencies.
type (

	// Stringer is implemented by any value that has a String method,
	// which defines the ``native'' format for that value.
	// The String method is used to print values passed as an operand
	// to any format that accepts a string or to an unformatted printer
	// such as Print.
	//
	// Ref: Standard Library fmt package
	Stringer interface {
		String() string
	}

	// State represents the printer state passed to custom formatters.
	// It provides access to the io.Writer interface plus information about
	// the flags and options for the operand's format specifier.
	//
	// Ref: Standard Library fmt package
	State interface {
		// Write is the function to call to emit formatted output to be printed.
		Write(b []byte) (n int, err error)
		// Width returns the value of the width option and whether it has been set.
		Width() (wid int, ok bool)
		// Precision returns the value of the precision option and whether it has been set.
		Precision() (prec int, ok bool)

		// Flag reports whether the flag c, a character, has been set.
		Flag(c int) bool
	}

	// Formatter is implemented by any value that has a Format method.
	// The implementation controls how State and rune are interpreted,
	// and may call Sprint(f) or Fprint(f) etc. to generate its output.
	//
	// Ref: Standard Library fmt package
	Formatter interface {
		Format(f State, verb rune)
	}

	// Writer is the interface that wraps the basic Write method.
	//
	// Write writes len(p) bytes from p to the underlying data stream.
	// It returns the number of bytes written from p (0 <= n <= len(p))
	// and any error encountered that caused the write to stop early.
	// Write must return a non-nil error if it returns n < len(p). Write
	// must not modify the slice data, even temporarily.
	//
	// Implementations must not retain p.
	//
	// Ref: Standard Library io package
	Writer interface {
		Write(p []byte) (n int, err error)
	}

	// StringWriter is the interface that wraps the WriteString method.
	//
	// WriteString writes the contents of the string s to w, which accepts a slice of bytes.
	// If w implements StringWriter, its WriteString method is invoked directly.
	// Otherwise, w.Write is called exactly once.
	//
	// Ref: Standard Library io package
	StringWriter interface {
		WriteString(s string) (n int, err error)
	}

	// A Locker represents an object that can be locked and unlocked.
	//	 type Locker interface {
	//	 	Lock()
	//	 	Unlock()
	//	 }
	//
	// Ref: Standard Library sync package
	Locker interface {
		Lock()
		Unlock()
	}

	// An implementation of sort.Interface can be sorted by the routines in
	// this package. The methods refer to elements of the underlying
	// collection by integer index.
	//
	// Ref: Standard Library sort package
	Sorter interface {
		// Len is the number of elements in the collection.
		Len() int

		// Less reports whether the element with index i
		// must sort before the element with index j.
		//
		// If both Less(i, j) and Less(j, i) are false,
		// then the elements at index i and j are considered equal.
		// Sort may place equal elements in any order in the final result,
		// while Stable preserves the original input order of equal elements.
		//
		// Less must describe a transitive ordering:
		//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
		//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
		//
		// Note that floating-point comparison (the < operator on float32 or float64 values)
		// is not a transitive ordering when not-a-number (NaN) values are involved.
		// See Float64Slice.Less for a correct implementation for floating-point values.
		Less(i, j int) bool

		// Swap swaps the elements with indexes i and j.
		Swap(i, j int)
	}
)
