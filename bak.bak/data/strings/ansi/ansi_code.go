// Copyright (c) 2021 Michael Treanor
// https://
// MIT License

package ansi

import (
	"bytes"
	"fmt"
	"strings"
)

type Any = interface{}

const (
	ansiEncodeBasicFMT string = "\x1b[%v;%v;%vm"
	prefix             string = "\x1b["
	fg                 string = "38;5;"
	bg                 string = "48;5;"
	suffix             string = "m"
	ansiSep            string = ";"

	fg2 = ansiSep + fg
	bg2 = ansiSep + bg
)

var (
	defaultAnsiCode = NewColor("2", "0", "1")
)

func ansiEncodeBasic(fg, bg, ef string) string {
	return fmt.Sprintf(ansiEncodeBasicFMT, fg, bg, ef)
}

func NewColor(foreground, background, effect string) string {
	return newColorConcat(foreground, background, effect)
}

func newColorConcat(foreground, background, effect string) string {
	//     "\x1b[        %d      ;	  38;5;       %d      ;       48;5;     %d         m"
	return prefix + effect + ansiSep + fg + foreground + ansiSep + bg + background + suffix
}

func newColorConcat2(foreground, background, effect string) string {
	//     "\x1b[     %d    ;38;5;     %d      ;48;5;     %d         m"
	return prefix + effect + fg2 + foreground + bg2 + background + suffix
}

const ansitemplate = "\x1b[000;38;5;000;48;5;000m"
const z = '0'

var a [24]byte

func NewColorMake(foreground, background, effect string) string {
	//     "\x1b[     000    ;38;5;      000      ;48;5;     000        m"
	//      0,1      2,3,4    5-10     11,12,13   14-19    20,21,22    23

	var b = a[:0]

	b = []byte(ansitemplate)

	fmt.Printf("a: %q\n", a)
	fmt.Printf("b: %q\n", b)

	fz := 3 - len(foreground)
	bz := 3 - len(background)
	ez := 3 - len(effect)

	for i := 0; i < fz; i++ {
		b[2+i] = '0'
	}

	fmt.Println("fz: ", fz)
	fmt.Println("bz: ", bz)
	fmt.Println("ez: ", ez)

	foreground = strings.Repeat("0", fz) + foreground
	background = strings.Repeat("0", bz) + background
	effect = strings.Repeat("0", ez) + effect

	fmt.Println("foreground: ", foreground)
	fmt.Println("background: ", background)
	fmt.Println("effect: ", effect)

	return ""
}

func newColorSB(foreground, background, effect string) string {
	sb := strings.Builder{}
	sb.WriteString(prefix)
	sb.WriteString(effect)
	sb.WriteString(ansiSep)
	sb.WriteString(fg)
	sb.WriteString(foreground)
	sb.WriteString(ansiSep)
	sb.WriteString(bg)
	sb.WriteString(background)
	sb.WriteString(suffix)
	return sb.String()
}

func newColorBB(foreground, background, effect string) string {
	sb := bytes.Buffer{}
	sb.WriteString(prefix)
	sb.WriteString(effect)
	sb.WriteString(ansiSep)
	sb.WriteString(fg)
	sb.WriteString(foreground)
	sb.WriteString(ansiSep)
	sb.WriteString(bg)
	sb.WriteString(background)
	sb.WriteString(suffix)
	return sb.String()
}

func newColorJoin(foreground, background, effect string) string {
	return strings.Join(
		[]string{prefix + effect, fg + foreground, bg + background + suffix},
		ansiSep,
	)
}

func newColorSprintf(foreground, background, effect string) string {
	return fmt.Sprintf(ansiEncodeBasicFMT, effect, foreground, background)
}
