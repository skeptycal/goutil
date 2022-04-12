package sort

// Convenience types for common cases

// StringSlice attaches the methods of Interface to []string, sorting in increasing order.
type StringSlice []string

func (x StringSlice) Len() int           { return len(x) }
func (x StringSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x StringSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Sort is a convenience method: x.Sort() calls Sort(x).
func (x StringSlice) Sort() { Sort(x) }

// Convenience wrappers for common cases

// Strings sorts a slice of strings in increasing order.
func Strings(x []string) { Sort(StringSlice(x)) }

// StringsAreSorted reports whether the slice x is sorted in increasing order.
func StringsAreSorted(x []string) bool { return IsSorted(StringSlice(x)) }
