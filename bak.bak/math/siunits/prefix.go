package siunits

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

func (si siMap) valid(key int) bool {
	if _, ok := si[key]; ok {
		return true
	} else {
		return false
	}
}

func (si siMap) get(key int) []string {
	if v, ok := si[key]; ok {
		return v
	}
	return defaultSIValue
}
