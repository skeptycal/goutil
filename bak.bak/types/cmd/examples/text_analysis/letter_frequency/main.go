package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"

	"github.com/skeptycal/goutil/types"
)

func GetSampleText() string {
	sampleFile := "romeo_and_juliet.txt"

	b, err := os.ReadFile(sampleFile)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

type ByteMap map[byte]int

func (b ByteMap) String() string {
	sb := strings.Builder{}
	defer sb.Reset()

	for k, v := range b {
		s := fmt.Sprintf("%c : %v\n", k, v)
		sb.WriteString(s)
	}

	return sb.String()
}

func main() {

	s := GetSampleText()

	var m ByteMap = ByteMap(types.Frequency(s))

	list := make([]byte, 0, len(m))

	for k := range m {
		if unicode.IsPrint(rune(k)) {
			list = append(list, k)
		}
	}

	sort.Slice(list, func(i, j int) bool { return list[i] < list[j] })

	for _, k := range list {
		fmt.Printf("%q : %v\n", k, m[k])
	}

	// fmt.Println(m)
}
