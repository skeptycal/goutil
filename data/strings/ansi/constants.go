// Copyright (c) 2020 Michael Treanor
// MIT License

// Package ansi provides an extensive set of ANSI escape codes
// for use in command line interfaces (CLIs) using industry standard
// well-documented terminal commands.
package ansi

// Format Strings 8 bit Ansi printf commands.
//
// ESC[⟨x⟩8:5:⟨n⟩m
//
// Select 8bit color
//
// n in [0..255]; 0-231 are colors; 232-255 are grayscale
//
// (x in [3, 4]); 3 = foreground; 4 = background
const (
	Fmt8bit   string = "\033[%d8;5;%dm"
	Fmt8bitFG string = "\033[38;5;%dm"
	Fmt8bitBG string = "\033[48;5;%dm"
)

// Format Strings - 3/4 bit Ansi printf commands.
//
// ESC[⟨e⟩;⟨x⟩⟨n⟩m
//
// Select 3/4 bit color
//
// n in [0..7]; basic colors
//
// (x in [3, 4]); 3 = foreground; 4 = background
//
// e can be any valid ANSI escape effect; these are common:
//  0 Normal            // may be left out to maintain current ANSI effect
//  1 Bold              // bold or increased intensity
//  2 Faint             // faint, decreased intensity or second color
//  3 Italics
//  4 Underline
//  7 Inverse
//  8 Conceal
//  9 Strikeout
//
const (
	FmtANSI           string = "\033[%dm"
	FmtANSIFG         string = "\033[3%dm"
	FmtANSIBG         string = "\033[4%dm"
	FmtANSIWithEffect string = "\033[%d;%dm"
	FmtANSIBright     string = "\033[1;%dm"
	FmtANSIDim        string = "\033[2;%dm"
)

// Format Strings 24 bit Ansi printf commands.
//
// ESC[⟨x⟩8;2;⟨R⟩;⟨G⟩;⟨B⟩m
//
// Select RGB color
//
// R, G, B in [0..255]
//
// (x in [3, 4]); 3 = foreground; 4 = background
const (
	Fmt24bit   string = "\033[%d8;2;%d;%d;%dm"
	Fmt24bitFG string = "\033[38;2;%d;%d;%dm"
	Fmt24bitBG string = "\033[48;2;%d;%d;%dm"
)

// premade ANSI basic 3 bit color codes
//
// Reference: https://en.wikipedia.org/wiki/ANSI_escape_code
const (
	BlackText          string = "\033[30m"
	RedText            string = "\033[31m"
	GreenText          string = "\033[32m"
	YellowText         string = "\033[33m"
	BlueText           string = "\033[34m"
	PurpleText         string = "\033[35m"
	CyanText           string = "\033[36m"
	WhiteText          string = "\033[37m"
	DefaultColorText   string = "\033[39m" // Normal foreground color
	BgBlackText        string = "\033[40m"
	BgRedText          string = "\033[41m"
	BgGreenText        string = "\033[42m"
	BgYellowText       string = "\033[43m"
	BgBlueText         string = "\033[44m"
	BgPurpleText       string = "\033[45m"
	BgCyanText         string = "\033[46m"
	BgWhiteText        string = "\033[47m"
	BhDefaultColorText string = "\033[49m" // Normal background color
)

// premade bold and dim ANSI text colors
const (
	BoldText         string = "\033[1m"
	BoldBlackText    string = "\033[1;30m"
	BoldRedText      string = "\033[1;31m"
	BoldGreenText    string = "\033[1;32m"
	BoldYellowText   string = "\033[1;33m"
	BoldBlueText     string = "\033[1;34m"
	BoldMagentaText  string = "\033[1;35m"
	BoldCyanText     string = "\033[1;36m"
	BoldWhiteText    string = "\033[1;37m"
	FaintText        string = "\033[2m"
	FaintBlackText   string = "\033[2;30m"
	FaintRedText     string = "\033[2;31m"
	FaintGreenText   string = "\033[2;32m"
	FaintYellowText  string = "\033[2;33m"
	FaintBlueText    string = "\033[2;34m"
	FaintMagentaText string = "\033[2;35m"
	FaintCyanText    string = "\033[2;36m"
	FaintWhiteText   string = "\033[2;37m"
)

// ANSI escape codes for text effects
//
// These are the most commonly used. ANSI codes above 9 are very
// rare and many are not fully implemented.
//  Normal      = 0
//  Bold        = 1
//  Faint       = 2
//  Italics     = 3
//  Underline   = 4
//  Inverse     = 7
//  Conceal     = 8
//  Strikeout   = 9
//
const (
	Normal byte = iota
	Bold        // bold or increased intensity
	Faint       // faint, decreased intensity or second color
	Italics
	Underline
	Blink
	FastBlink
	Inverse
	Conceal
	Strikeout
	// ANSI codes above 9 are very rare and many are not fully implemented.
	PrimaryFont
	AltFont1
	AltFont2
	AltFont3
	AltFont4
	AltFont5
	AltFont6
	AltFont7
	AltFont8
	AltFont9
	Gothic // fraktur
	DoubleUnderline
	NormalColor // normal color or normal intensity (neither bold nor faint)
	NotItalics  // not italicized, not fraktur
	NotUnderlined
	Steady     // not Blink or FastBlink
	Reserved26 // reserved for proportional spacing as specified in CCITT Recommendation T.61
	NotInverse // Positive
	NotHidden  // Revealed
	NotStrikeout
	Black
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	SetForeground     // Next arguments are 5;n or 2;r;g;b, see below
	DefaultForeground // default display color (implementation-defined)
	BlackBackground
	RedBackground
	GreenBackground
	YellowBackground
	BlueBackground
	MagentaBackground
	CyanBackground
	WhiteBackground
	SetBackground              // Next arguments are 5;n or 2;r;g;b, see below
	DefaultBackground          // default background color (implementation-defined)
	DisableProportionalSpacing // reserved for cancelling the effect of parameter value 26
	Framed
	Encircled
	Overlined
	NotFramed // NotEncircled
	NotOverlined
	Reserved56
	Reserved57
	SetUnderlineColor // Next arguments are 5;n or 2;r;g;b, see below
	DefaultUnderlineColor
	IdeogramUnderline       // ideogram underline or right side line
	IdeogramDoubleUnderline // ideogram double underline or double line on the right side
	IdeogramOverline        // ideogram overline or left side line
	IdeogramDoubleOverline  // ideogram double overline or double line on the left side
	IdeogramStress          // ideogram stress marking
	IdeogramCancel          // reset the effects of all of 60–64
	Superscript             = 73
	Subscript               = 74
)
