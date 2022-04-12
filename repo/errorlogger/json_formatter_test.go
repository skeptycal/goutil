package errorlogger

import (
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
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

var (
	// smallFields is a small size data set for benchmarking
	smallFields = Fields{
		"foo":   "bar",
		"baz":   "qux",
		"one":   "two",
		"three": "four",
	}

	// largeFields is a large size data set for benchmarking
	largeFields = Fields{
		"foo":       "bar",
		"baz":       "qux",
		"one":       "two",
		"three":     "four",
		"five":      "six",
		"seven":     "eight",
		"nine":      "ten",
		"eleven":    "twelve",
		"thirteen":  "fourteen",
		"fifteen":   "sixteen",
		"seventeen": "eighteen",
		"nineteen":  "twenty",
		"a":         "b",
		"c":         "d",
		"e":         "f",
		"g":         "h",
		"i":         "j",
		"k":         "l",
		"m":         "n",
		"o":         "p",
		"q":         "r",
		"s":         "t",
		"u":         "v",
		"w":         "x",
		"y":         "z",
		"this":      "will",
		"make":      "thirty",
		"entries":   "yeah",
	}

	errorFields = Fields{
		"foo": fmt.Errorf("bar"),
		"baz": fmt.Errorf("qux"),
	}

	sampleTextFormatter        = &TextFormatter{logrus.TextFormatter{DisableColors: true}}
	sampleColoredTextFormatter = &TextFormatter{logrus.TextFormatter{ForceColors: true}}
	sampleJSONFormatter        = &JSONFormatter{logrus.JSONFormatter{PrettyPrint: false}}
	sampleJSONPrettyFormatter  = &JSONFormatter{logrus.JSONFormatter{PrettyPrint: true}}

	formatterTests = []struct {
		name   string
		f      Formatter
		fields Fields
	}{
		{"ErrorTextFormatter", sampleTextFormatter, errorFields},
		{"SmallTextFormatter", sampleTextFormatter, smallFields},
		{"LargeTextFormatter", sampleTextFormatter, largeFields},
		{"SmallColoredTextFormatter", sampleTextFormatter, smallFields},
		{"LargeColoredTextFormatter", sampleTextFormatter, largeFields},
		{"SmallJSONFormatter", sampleJSONFormatter, smallFields},
		{"LargeJSONFormatter", sampleJSONFormatter, largeFields},
		{"SmallJSONPrettyFormatter", sampleJSONFormatter, smallFields},
		{"LargeJSONPrettyFormatter", sampleJSONFormatter, largeFields},
	}
)

func BenchmarkFormatters(b *testing.B) {
	for _, bb := range formatterTests {
		doBenchmark(b, bb.name, bb.f, bb.fields)
	}
}

func doBenchmark(b *testing.B, name string, formatter Formatter, fields Fields) {
	logger := logrus.New()

	entry := &Entry{
		Time:    time.Time{},
		Level:   InfoLevel,
		Message: "message",
		Data:    fields,
		Logger:  logger,
	}
	var d []byte
	var err error
	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d, err = formatter.Format(entry)
			if err != nil {
				b.Fatal(err)
			}
			b.SetBytes(int64(len(d)))
		}
	})
}

func TestNewJSONFormatter(t *testing.T) {
	tests := []struct {
		name   string
		pretty bool
		want   Formatter
	}{
		{"new default JSON formatter", true, &JSONFormatter{logrus.JSONFormatter{PrettyPrint: true}}},
	}
	log.SetOutput(ioutil.Discard)
	for _, tt := range tests {
		fakeFunc := func(*runtime.Frame) (function string, file string) { return "", "" }
		fakeMap := logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg:   "@message",
			logrus.FieldKeyFunc:  "@caller",
		}

		t.Run(tt.name, func(t *testing.T) {
			got := NewJSONFormatter(tt.pretty)

			got.SetPrettyPrint(false)
			got.SetPrettyPrint(true)
			got.SetPrettyPrint(tt.pretty)
			got.SetTimestampFormat(time.UnixDate) // UnixDate is just something to test with
			got.SetTimestampFormat(time.RFC3339)  // RFC3339 is the default
			got.SetDisableTimeStamp(true)
			got.SetDisableTimeStamp(false)
			got.SetDisableHTMLEscape(false)
			got.SetDisableHTMLEscape(true)
			got.SetDataKey("key")
			got.SetFieldMap(fakeMap)
			got.SetCallerPrettyfier(fakeFunc)
			f := got.Formatter()
			// f is already a Formatter by definition, so this should never fail ...
			if _, ok := f.(Formatter); !ok {
				t.Errorf("JSONFormatter.Formatter() does not implement Formatter interface: %v(%T)", f, f)
			}
		})
	}
}
