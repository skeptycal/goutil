package errorlogger

import (
	"runtime"

	"github.com/sirupsen/logrus"
)

// DefaultTextFormatter is the default log formatter. Use
//  Log.SetText()
// or
//  Log.SetFormatter(DefaultTextFormatter)
// to return to default text formatting of logs.
//
// To change to another logrus formatter, use
//  Log.SetFormatter(myFormatter)
//
// Reference: https://pkg.go.dev/github.com/sirupsen/logrus#TextFormatter
var DefaultTextFormatter Formatter = NewTextFormatter()

// TextFormatter formats logs into text.
// Note: this is a logrus type with marginally useful
// utilities and may become a type alias in the future.
/*
	type TextFormatter struct {
		// Set to true to bypass checking for a TTY before outputting colors.
		ForceColors bool

		// Force disabling colors.
		DisableColors bool

		// Force quoting of all values
		ForceQuote bool

		// DisableQuote disables quoting for all values.
		// DisableQuote will have a lower priority than ForceQuote.
		// If both of them are set to true, quote will be forced on all values.
		DisableQuote bool

		// Override coloring based on CLICOLOR and CLICOLOR_FORCE.
		// Reference: https://bixense.com/clicolors/
		EnvironmentOverrideColors bool

		// Disable timestamp logging. useful when output is redirected to logging
		// system that already adds timestamps.
		DisableTimestamp bool

		// Enable logging the full timestamp when a TTY is attached instead of just
		// the time passed since beginning of execution.
		FullTimestamp bool

		// TimestampFormat to use for display when a full timestamp is printed.
		// The format to use is the same than for time.Format or time.Parse from the standard
		// library.
		// The standard Library already provides a set of predefined format.
		TimestampFormat string

		// The fields are sorted by default for a consistent output. For applications
		// that log extremely frequently and don't use the JSON formatter this may not
		// be desired.
		DisableSorting bool

		// The keys sorting function, when uninitialized it uses sort.Strings.
		SortingFunc func([]string)

		// Disables the truncation of the level text to 4 characters.
		DisableLevelTruncation bool

		// PadLevelText Adds padding the level text so that all the levels output at the same length
		// PadLevelText is a superset of the DisableLevelTruncation option
		PadLevelText bool

		// QuoteEmptyFields will wrap empty fields in quotes if true
		QuoteEmptyFields bool

		// Whether the logger's out is to a terminal
		isTerminal bool

		// FieldMap allows users to customize the names of keys for default fields.
		// As an example:
		// formatter := &TextFormatter{
		//     FieldMap: FieldMap{
		//         FieldKeyTime:  "@timestamp",
		//         FieldKeyLevel: "@level",
		//         FieldKeyMsg:   "@message"}}
		FieldMap FieldMap

		// CallerPrettyfier can be set by the user to modify the content
		// of the function and file keys in the data when ReportCaller is
		// activated. If any of the returned value is the empty string the
		// corresponding key will be removed from fields.
		CallerPrettyfier func(*runtime.Frame) (function string, file string)

		terminalInitOnce sync.Once

		// The max length of the level text, generated dynamically on init
		levelTextMaxLength int
	}
*/
type TextFormatter struct {
	logrus.TextFormatter
}

// NewTextFormatter returns a new TextFormatter that
// is initialized and ready to use.
func NewTextFormatter() *TextFormatter {
	return &TextFormatter{}
}

// SetForceColors allows users to bypass checking for a TTY
// before outputting colors and forces color output.
func (f *TextFormatter) SetForceColors(yesno bool) {
	f.ForceColors = yesno
}

// SetDisableColors allows users to disable colors.
func (f *TextFormatter) SetDisableColors(yesno bool) {
	f.DisableColors = yesno
}

// SetForceQuote allows users to force quoting of all values.
func (f *TextFormatter) SetForceQuote(yesno bool) {
	f.ForceQuote = yesno
}

// SetDisableQuote allows users to disable quoting for all values.
// It has a lower priority than SetForceQuote, i.e. if both of them
// are set to true, quotes will be forced on for all values.
func (f *TextFormatter) SetDisableQuote(yesno bool) {
	f.DisableQuote = yesno
}

// SetEnvironmentOverrideColors allows users to override coloring based
// on CLICOLOR and CLICOLOR_FORCE. - https://bixense.com/clicolors/
func (f *TextFormatter) SetEnvironmentOverrideColors(yesno bool) {
	f.EnvironmentOverrideColors = yesno
}

// SetDisableTimeStamp allows users to disable automatic timestamp logging.
// Useful when output is redirected to logging systems that already
// add timestamps.
func (f *TextFormatter) SetDisableTimeStamp(yesno bool) {
	f.DisableTimestamp = yesno
}

// SetFullTimeStamp allows users to enable logging the full timestamp
// when a TTY is attached instead of just the time passed since beginning
// of execution.
func (f *TextFormatter) SetFullTimeStamp(yesno bool) {
	f.FullTimestamp = yesno
}

// SetTimestampFormat sets the format for display when a full
// timestamp is printed. The format to use is the same than for
// time.Format or time.Parse from the standard library.
// The standard Library already provides a set of predefined formats.
// The recommended and default format is time.RFC3339.
func (f *TextFormatter) SetTimestampFormat(fmt string) {
	f.TimestampFormat = fmt
}

// SetDisableSorting allows users to disable the default behavior
// of sorting of fields by default for a consistent output. For
// applications that log extremely frequently and don't use the
// JSON formatter this may not be desired.
func (f *TextFormatter) SetDisableSorting(yesno bool) {
	f.DisableSorting = yesno
}

// SetSortingFunc allows users to set the keys sorting function.
// The default is sort.Strings.
func (f *TextFormatter) SetSortingFunc(fn func([]string)) {
	f.SortingFunc = fn
}

// SetDisableLevelTruncation allows users to disable the truncation of the level text to 4 characters.
func (f *TextFormatter) SetDisableLevelTruncation(yesno bool) {
	f.DisableLevelTruncation = yesno
}

// SetPadLevelText allows users to enable the addition of padding
// to the level text so that all the levels output at the same length
// PadLevelText is a superset of the DisableLevelTruncation option
func (f *TextFormatter) SetPadLevelText(yesno bool) {
	f.PadLevelText = yesno
}

// SetQuoteEmptyFields allows users to enable the wrapping of empty
// fields in quotes.
func (f *TextFormatter) SetQuoteEmptyFields(yesno bool) {
	f.QuoteEmptyFields = yesno
}

// SetFieldMap allows users to customize the names of keys
// for default fields.
// For example:
//  formatter := &TextFormatter{
//   	FieldMap: FieldMap{
// 		 FieldKeyTime:  "@timestamp",
// 		 FieldKeyLevel: "@level",
// 		 FieldKeyMsg:   "@message",
// 		 FieldKeyFunc:  "@caller",
//    },
//  }
func (f *TextFormatter) SetFieldMap(m logrus.FieldMap) {
	f.FieldMap = m
}

// SetCallerPrettyfier sets the user option to modify the content
// of the function and file keys in the data when ReportCaller is
// activated. If any of the returned values is the empty string the
// corresponding key will be removed from fields.
func (f *TextFormatter) SetCallerPrettyfier(fn func(*runtime.Frame) (function string, file string)) {
	f.CallerPrettyfier = fn
}
