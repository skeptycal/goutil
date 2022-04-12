package anansi

import (
	"os"

	"github.com/fatih/color"
	"github.com/skeptycal/goutil/repo/errorlogger"
	// . "github.com/fatih/color"
)

var (
	Log = errorlogger.Log
)

// NoColor defines if the output is colorized or not. It's dynamically set to false or true based on the stdout's file descriptor referring to a terminal or not. It's also set to true if the NO_COLOR environment variable is set (regardless of its value). This is a global option and affects all colors. For more control over each color block use the methods DisableColor() individually.
var NoColor = color.NoColor

// ANSI color printers
var (
	defaultStdout = os.Stdout

	Reset       = NewColor(ResetCode)
	Canary      = NewColor(FgHiYellow).Add(Bold)
	Green       = NewColor(FgHiGreen)
	Whataburger = NewColor(Attribute(color.FgRed))

	ATTN = Canary.Add(BgRed)
	INFO = Green.Add(Bold).Add(Italic).Add(BgHiBlue)
)

type (

	// Attribute defines a single SGR Code
	//
	// Reference: http://github.com/fatih/color```` (MIT License)
	Attribute = color.Attribute

	// Color defines a custom color object which is defined by SGR parameters.
	//
	// Reference: http://github.com/fatih/color (MIT License)
	Color = color.Color
)

// NewColor returns a newly created color object.
func NewColor(value ...Attribute) *Color {
	return color.New(value...)
}

// Attn outputs INFO level information to the specified
// log io.Writer. If the writer supports ANSI escape
// sequences, the text will be written as yellow text on
// a red background.
func Attn(format string, args ...interface{}) {
	Log.Warn(ATTN.Sprintf(format, args...))
}

// Info outputs INFO level information to the specified
// log io.Writer. If the writer supports ANSI escape
// sequences, the text will be written as yellow text on
// a red background.
func Info(format string, args ...interface{}) {
	Log.Info(INFO.Sprintf(format, args...))
}

func Blue(format string, args ...interface{}) (n int, err error) {
	return color.New(color.FgBlue).Fprintf(defaultStdout, format, args...)
}

// constants from https://github.com/fatih/color

// Base attributes
const (
	ResetCode Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground text colors
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors
const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity text colors
const (
	BgHiBlack Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)
