package benchmark

import "fmt"

// itemList is a collection of test data for testFuncs.
type testData struct {
	name    string
	input   Any
	want    Any
	wantErr bool
}

// String returns a string representation of the test data
// formatted using the default format string:
//  fmtTestInfo
func (t *testData) String() string {
	return fmt.Sprintf(fmtTestInfo, t.name, t.input, t.want, t.wantErr)
}
