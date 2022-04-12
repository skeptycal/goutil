package benchmark

// Benchmarks is a collection of test sets to run
// on the entire code base. This is the top level structure.
type Benchmarks struct {
	name   string
	note   string
	config *ConfigSettings
	tests  []BenchmarkTestSet
}

// BenchmarkTestSet is a collection of benchmark tests
// and configuration information. Allows each benchmark
// to have its own configuration settings.
//
// Use <nil> for config to use parent configuration.
type BenchmarkTestSet struct {
	name   string
	config *ConfigSettings
	list   testList
}

// testList is a collection of individual tests to run
// on a specified function.
type testList struct {
	name     string
	args     []testData
	funcList []testFuncs
}
