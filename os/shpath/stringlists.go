package shpath

func RemoveOrdered(s []string, pos int) []string {
	return append(s[:pos], s[pos+1:]...)
}

func RemoveUnOrdered(s []string, n int) []string {
	s[n] = s[len(s)-1]
	s[len(s)-1] = "" // clear the last element before removing it ... maybe helps with GC?
	return s[:len(s)-1]
}

func Append(list []string, s string) []string {
	return append(list, s)
}

func Insert(list []string, s string, pos int) []string {

	temp := make([]string, len(list)+1) // preallocating is faster on most benchmarks

	copy(temp, list[:pos]) // copy is faster than append on most benchmarks
	temp = append(temp, s)
	copy(temp, list[pos+1:])

	list = nil // maybe helps GC?

	return temp
}
