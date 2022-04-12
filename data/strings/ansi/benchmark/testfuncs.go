package benchmark

// testFuncs a test function to run with
// with testData values.
type testFuncs struct {
	name string
	fn   func(...Any) Any
}
