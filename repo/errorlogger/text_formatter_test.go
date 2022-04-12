package errorlogger

import (
	"runtime"
	"sort"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

//////////////////////// Benchmarks
// inspired by (and borrowed from)
// Reference: http://github.com/sirupsen/logrus
/*
ErrorTextFormatter-8         	  574843      2099 ns/op    31.92 MB/s     536 B/op      16 allocs/op
SmallTextFormatter-8         	  616621      1957 ns/op    43.94 MB/s     592 B/op      14 allocs/op
LargeTextFormatter-8         	   88616     13626 ns/op    20.99 MB/s    6549 B/op      20 allocs/op
SmallColoredTextFormatter-8  	  615572      1955 ns/op    43.98 MB/s     592 B/op      14 allocs/op
LargeColoredTextFormatter-8  	   86119     13435 ns/op    21.29 MB/s    6550 B/op      20 allocs/op
SmallJSONFormatter-8         	  310737      3869 ns/op    29.46 MB/s    1528 B/op      30 allocs/op
LargeJSONFormatter-8         	   69032     17372 ns/op    23.60 MB/s    6226 B/op      78 allocs/op
SmallJSONPrettyFormatter-8   	  310528      3880 ns/op    29.38 MB/s    1528 B/op      30 allocs/op
LargeJSONPrettyFormatter-8   	   69073     17311 ns/op    23.68 MB/s    6226 B/op      78 allocs/op
*/

func TestNewTextFormatter(t *testing.T) {
	tests := []struct {
		name string
		want Formatter
	}{
		{"new default JSON formatter", &TextFormatter{logrus.TextFormatter{}}},
	}
	for _, tt := range tests {

		fakeFunc := func(*runtime.Frame) (function string, file string) { return "", "" }
		fakeMap := logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg:   "@message",
			logrus.FieldKeyFunc:  "@caller",
		}

		t.Run(tt.name, func(_ *testing.T) {

			got := NewTextFormatter()

			got.SetForceColors(false)
			got.SetForceColors(true)
			got.SetDisableColors(true)
			got.SetDisableColors(false)
			got.SetForceQuote(true)
			got.SetForceQuote(false)
			got.SetDisableQuote(false)
			got.SetDisableQuote(true)
			got.SetEnvironmentOverrideColors(true)
			got.SetEnvironmentOverrideColors(false)
			got.SetDisableTimeStamp(true)
			got.SetDisableTimeStamp(false)
			got.SetFullTimeStamp(true)
			got.SetFullTimeStamp(false)
			got.SetTimestampFormat(time.UnixDate)
			got.SetTimestampFormat(time.RFC3339)
			got.SetDisableSorting(false)
			got.SetDisableSorting(true)
			got.SetSortingFunc(nil)
			got.SetSortingFunc(func(s []string) { sort.Strings(s) })
			got.SetDisableLevelTruncation(false)
			got.SetDisableLevelTruncation(true)
			got.SetPadLevelText(true)
			got.SetPadLevelText(true)
			got.SetQuoteEmptyFields(true)
			got.SetQuoteEmptyFields(true)
			got.SetFieldMap(nil)
			got.SetFieldMap(fakeMap)
			got.SetCallerPrettyfier(fakeFunc)
		})
	}
}
