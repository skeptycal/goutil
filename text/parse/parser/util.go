package parser

import (
	"strings"
	"unicode"
)

func fold(in string) string {
	sb := strings.Builder{}
	defer sb.Reset()

	for _, r := range in {
		sb.WriteRune(unicode.SimpleFold(r))
	}
	return sb.String()
}

func spacer(in, repl string, firstLower, titleCase bool) (retval string) {

	if titleCase {
		in = strings.ToTitle(in)
	}

	if firstLower {
		retval = strings.ToLower(in[:1])
	} else {
		retval = in[1:]
	}

	retval += strings.ReplaceAll(in[1:], " ", repl)

	return retval
}

func ToSnake(in string) string {
	return spacer(strings.ToLower(in), "_", true, false)
}

func ToPascal(in string) string {
	return ""
}
