package errorlogger

import (
	"runtime"

	"github.com/sirupsen/logrus"
)

// JSONFormatter formats logs into parsable json.
// It is composed of logrus.JSONFormatter with additional
// formatting methods.
type JSONFormatter struct{ logrus.JSONFormatter }

// NewJSONFormatter returns a new Formatter that
// is initialized and ready to use.
//
// For pretty printing, set pretty == true.
func NewJSONFormatter(pretty bool) *JSONFormatter {
	f := &JSONFormatter{logrus.JSONFormatter{}}
	f.SetPrettyPrint(pretty)
	return f
}

func (f *JSONFormatter) Formatter() Formatter {
	return &f.JSONFormatter
}

// SetTimestampFormat sets the format used for marshaling timestamps.
// The format to use is the same as for time.Format or time.Parse
// from the standard library.
// The standard Library already provides a set of predefined formats.
// The recommended and default format is RFC3339.
func (f *JSONFormatter) SetTimestampFormat(fmt string) {
	f.TimestampFormat = fmt
}

// SetDisableTimeStamp allows disabling automatic timestamps in output
func (f *JSONFormatter) SetDisableTimeStamp(yesno bool) {
	f.DisableTimestamp = yesno
}

// SetDisableHTMLEscape allows disabling html escaping in output
func (f *JSONFormatter) SetDisableHTMLEscape(yesno bool) {
	f.DisableHTMLEscape = yesno
}

// SetDataKey allows users to put all the log entry parameters
// into a nested dictionary at a given key.
func (f *JSONFormatter) SetDataKey(key string) {
	f.DataKey = key
}

// SetFieldMap allows users to customize the names of keys
// for default fields.
// For example:
//  formatter := &JSONFormatter{
//   	FieldMap: FieldMap{
// 		 FieldKeyTime:  "@timestamp",
// 		 FieldKeyLevel: "@level",
// 		 FieldKeyMsg:   "@message",
// 		 FieldKeyFunc:  "@caller",
//    },
//  }
func (f *JSONFormatter) SetFieldMap(m logrus.FieldMap) {
	f.FieldMap = m
}

// SetCallerPrettyfier sets the user option to modify the content
// of the function and file keys in the json data when ReportCaller is
// activated. If any of the returned values is the empty string the
// corresponding key will be removed from json fields.
func (f *JSONFormatter) SetCallerPrettyfier(fn func(*runtime.Frame) (function string, file string)) {
	f.CallerPrettyfier = fn
}

// SetPrettyPrint set to true will indent all json logs
func (f *JSONFormatter) SetPrettyPrint(pretty bool) {
	f.PrettyPrint = pretty
}
