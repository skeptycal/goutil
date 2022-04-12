package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Maximum File Size by Operating System
/*
File system	Maximum size
	APFS	 8 EB
	exFAT	16 EB
	FAT12	16 MB (4 KB clusters) or 32 MB (8 KB clusters)
	FAT16B	 2 GB (without LFS) or 4 GB (with LFS)
	FAT32	 4 GB
	HFS		 2 GB
	HFS+	 8 EB
	HPFS	 2 GB
	NTFS	16 EB

	Prefix meanings have been notoriously inconsistent throughout
	the history of the computer industry: https://en.wikipedia.org/wiki/Binary_prefix

	For clarity (and consistency with my science background) I use the SI system:
	https://en.wikipedia.org/wiki/International_System_of_Units

*/
const (
	pattern         = "."
	maxPrefixLength = 18
	maxPrefix       = "Gb"
)

type (

	// SIPrefix creates an interface for accessing names and symbols for SI unit prefixes.
	//
	// The SI prefixes are metric prefixes that were standardized for use in the International System of Units (SI) by the International Bureau of Weights and Measures (BIPM) in resolutions dating from 1960 to 1991. Since 2009, they have formed part of the International System of Quantities. They are also used in the Unified Code for Units of Measure (UCUM)
	//
	// A metric prefix is a unit prefix that precedes a basic unit of measure to indicate a multiple or submultiple of the unit. All metric prefixes used today are decadic. Each prefix has a unique symbol that is prepended to any unit symbol. The prefix kilo-, for example, may be added to gram to indicate multiplication by one thousand: one kilogram is equal to one thousand grams. The prefix milli-, likewise, may be added to metre to indicate division by one thousand; one millimetre is equal to one thousandth of a metre.
	//
	// Decimal multiplicative prefixes have been a feature of all forms of the metric system, with six of these dating back to the system's introduction in the 1790s. Metric prefixes have also been used with some non-metric units. The SI prefixes are metric prefixes that were standardized for use in the International System of Units (SI) by the International Bureau of Weights and Measures (BIPM) in resolutions dating from 1960 to 1991. Since 2009, they have formed part of the International System of Quantities. They are also used in the Unified Code for Units of Measure (UCUM)
	//
	// Each prefix name has a symbol that is used in combination with the symbols for units of measure. For example, the symbol for kilo- is k, and is used to produce km, kg, and kW, which are the SI symbols for kilometre, kilogram, and kilowatt, respectively. Except for the early prefixes of kilo-, hecto-, and deca-, the symbols for the multiplicative prefixes are uppercase letters, and those for the fractional prefixes are lowercase letters.[2] There is a Unicode symbol for micro µ for use if the Greek letter μ is unavailable.[Note 1] When both are unavailable, the visually similar lowercase Latin letter u is commonly used instead. SI unit symbols are never italicised.
	//
	// References:
	//
	// Metric System (Wikipedia): https://en.wikipedia.org/wiki/Metric_prefix
	//
	// BIPM: https://www.bipm.org/en/home
	//
	// SI Brochure: The International System of Units (SI): https://www.bipm.org/en/publications/si-brochure (english)
	//
	// The International System of Units - 9th Edition (original French) https://www.bipm.org/documents/20126/41483022/SI-Brochure-9.pdf
	// The International System of Units - 9th Edition (English) https://www.bipm.org/documents/20126/41483022/SI-Brochure-9-EN.pdf
	SIPrefix interface {
		Name(i int) string
		Symbol(i int) string
	}

	siMap map[int][]string
)

const (
	defaultSIUnitName   = ""
	defaultSISymbolName = ""
)

func newSI() SIPrefix {
	return &siDefaultMap
}

var (
	SI             = newSI()
	defaultSIValue = []string{defaultSIUnitName, defaultSISymbolName}

	siDefaultMap siMap = siMap{
		0:   []string{"", ""},
		1:   []string{"da", "deca"}, // adopted 1795
		2:   []string{"h", "hecto"}, // adopted 1795
		3:   []string{"k", "kilo"},  // adopted 1795
		6:   []string{"M", "mega"},  // adopted 1873
		9:   []string{"G", "giga"},  // adopted 1960
		12:  []string{"T", "tera"},  // adopted 1960
		15:  []string{"P", "peta"},  // adopted 1975
		18:  []string{"E", "exa"},   // adopted 1975
		21:  []string{"Z", "zetta"}, // adopted 1991
		24:  []string{"Y", "yotta"}, // adopted 1991
		-1:  []string{"d", "deci"},  // adopted 1795
		-2:  []string{"c", "centi"}, // adopted 1795
		-3:  []string{"m", "milli"}, // adopted 1795
		-6:  []string{"µ", "micro"}, // adopted 1873 // µ (micro) or μ (Greek) or u (Latin)
		-9:  []string{"n", "nano"},  // adopted 1960
		-12: []string{"p", "pico"},  // adopted 1960
		-15: []string{"f", "femto"}, // adopted 1964
		-18: []string{"a", "atto"},  // adopted 1964
		-21: []string{"z", "zepto"}, // adopted 1991
		-24: []string{"y", "yocto"}, // adopted 1991
	}
)

func (si siMap) Symbol(key int) string { return si.get(key)[0] }
func (si siMap) Name(key int) string   { return si.get(key)[1] }

func (si siMap) get(key int) []string {
	if v, ok := si[key]; ok {
		return v
	}
	return defaultSIValue
}

func main() {

	// As of Go 1.16, os.ReadDir is a more efficient and correct choice: it returns a list of fs.DirEntry instead of fs.FileInfo, and it returns partial results in the case of an error midway through reading a directory.
	files, err := os.ReadDir(pattern)
	// files, err = ioutil.ReadDir(pattern)
	if err != nil {
		log.Fatal(err)
	}

	var fs fileSize
	for _, f := range files {
		fi, _ := f.Info()
		fs.int64 = fi.Size()
		fmt.Printf("%s %6v %s\n", f.Type(), fs, f.Name())

	}
}

type fileSize struct{ int64 }

func (f fileSize) Len() int { return len(f.ToString()) }

// 2
// 		return v[:w-9] + " GB"
// 	}
// 	if w > 6 {
// 		return v[:w-6] + " MB"
// 	}
// 	if w > 3 {
// 		return v[:w-3] + " KB"
// 	}
// 	return v + "   "
// }

func (f fileSize) ToString() string {
	return strconv.FormatInt(f.int64, 10)
}

func (f fileSize) String() string {

	v := strconv.FormatInt(f.int64, 10)
	w := len(v)

	if w > 9 {
		return v[:w-9] + " GB"
	}
	if w > 6 {
		return v[:w-6] + " MB"
	}
	if w > 3 {
		return v[:w-3] + " KB"
	}
	return v + "   "
}

/// ********************* not needed for this

type (
	Any interface{}

	// alternate version for performance profiling
	prefixSI struct {
		locked   bool
		prefixes map[int]string
	}
)

const SISEP = ":"

func parseSymbol(v string) []string {
	return strings.Split(v, SISEP)
}

func (si *prefixSI) Name(key int) Any {
	if v, err := si.Get(key); err == nil {
		return strings.Split(v.(string), SISEP)[0]
	}
	return nil
}

func (si *prefixSI) Symbol(key int) Any {
	if v, err := si.Get(key); err == nil {
		return strings.Split(v.(string), SISEP)[1]
	}
	return nil
}

func (si *prefixSI) Get(key Any) (Any, error) {
	if err := si.keyGuard(key); err != nil {
		return "", err
	}
	if v, ok := si.prefixes[key.(int)]; ok {
		return v, nil
	}
	return "", fmt.Errorf("key not found: %v", key)
}

func (si *prefixSI) Set(key, value Any) error {
	if err := si.guard(key, value); err != nil {
		return err
	}

	if _, ok := si.prefixes[key.(int)]; ok {
		if si.locked {
			return fmt.Errorf("duplicate key not allowed (data locked): %v", key)
		}
	}

	si.prefixes[key.(int)] = value.(string)
	return nil
}

func (si *prefixSI) Add(base int, name, symbol string) {}

func (si *prefixSI) keyGuard(key Any) error {
	switch key.(type) {
	case int:
		return nil
	default:
		return fmt.Errorf("wrong key type: %v(%T)", key, key)
	}
}

func (si *prefixSI) valueGuard(key Any) error {
	switch key.(type) {
	case string:
		return nil
	default:
		return fmt.Errorf("wrong value type: %v(%T)", key, key)
	}
}

func (si *prefixSI) guard(key, value Any) error {

	if err := si.keyGuard(key); err == nil {
		if _, ok := si.prefixes[key.(int)]; ok {
			return nil
		}
		return fmt.Errorf("key not found: %v", key)
	}
	err := si.keyGuard(key)
	if err != nil {
		return err
	}
	return si.valueGuard(value)
}
