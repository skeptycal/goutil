package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
	"
)

const (
	NL = "\n"
)

var (
	c                       string   = ansi.NewColor("172", "0", "1")
	reset                   string   = ansi.NewColor("7", "0", "0")
	ignorePrefixList        []string = []string{"//"}
	defaultPrintLineNumbers bool     = false
	defaultDoSort           bool     = true
	defaultBorderChar       string   = "="
)

func main() {
	fileName := "./sample/ansi_code.go"

	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	lines := GetFileLines(fileName, ignorePrefixList)

	if len(lines) == 0 {
		log.Fatal("list contains no lines of code.")
	}

	PrintList("Types:", FindTypes(lines), defaultPrintLineNumbers, defaultDoSort)
	PrintList("Functions:", FindFuncs(lines), defaultPrintLineNumbers, defaultDoSort)
}

func cPrint(args ...any) (n int, err error) {
	s := fmt.Sprint(args...)
	return fmt.Printf("%s%s%s\n", c, s, reset)
}

func formatHeader(borderChar string, width int, args ...any) string {
	s := fmt.Sprint(args...)
	if borderChar == "" {
		borderChar = defaultBorderChar
	}

	var leftborder, rightborder int
	border := width - len(s)
	if border < 2 {
		leftborder = 2
		rightborder = 0
	} else {
		leftborder = border / 2
		rightborder = border / 2
	}

	lb := strings.Repeat(borderChar, leftborder)
	rb := strings.Repeat(borderChar, rightborder)

	return fmt.Sprintf("%s %s %s", lb, s, rb)
}

func headerPrint(args ...any) (n int, err error) {
	// cPrint(formatHeader("=", 80, ""))
	n, err = cPrint(formatHeader("*", 80, args...))
	// cPrint(formatHeader("-", 80, ""))
	return
}

func footerPrint(args ...any) (n int, err error) {
	n, err = cPrint(formatHeader("*", 80, args...))
	return
}

// PrintList pretty prints a slice with defaults.
func PrintList(title string, list []string) {

}

func PPrintListWithOptions(title string, list []string, withNumbers, doSort, header, footer bool) {
	if doSort && !sort.StringsAreSorted(list) {
		sort.Strings(list)
	}
	headerPrint(title)

	if withNumbers {
		for i, line := range list {
			fmt.Printf(" %v%4d: %v%v\n", c, i, line, reset)
		}
	} else {
		for _, line := range list {
			fmt.Printf(" %v - %v%v\n", c, line, reset)
		}
	}
	// footerPrint()
	fmt.Println("")
}

func FindTypes(list []string) []string {
	var newlist []string
	for _, line := range list {
		if strings.HasPrefix(line, "type") {
			newlist = append(newlist, line)
		}
	}
	return newlist
}

func FindFuncs(list []string) []string {
	var newlist []string
	for _, line := range list {
		if strings.HasPrefix(line, "func") {
			if loc := strings.Index(line, "{"); loc > 1 {
				line = line[:loc-1]
			}
			newlist = append(newlist, normalizeWhitespace(line))
		}
	}
	return newlist
}

func GetFileLines(fileName string, ignorePrefix []string) []string {
	b, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), NL)
	newlist := []string{}

	for _, line := range lines {
		for _, ip := range ignorePrefix {
			if strings.HasPrefix(line, ip) {
				continue
			}
			newlist = append(newlist, line)
		}
	}

	return newlist
}

func normalizeWhitespace(in string) (out string) {
	return strings.Join(strings.Fields(in), " ")
}
