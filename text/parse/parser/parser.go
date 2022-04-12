package parser

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

func NewParser(r io.Reader) (Parser, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	p := parser{}
	p.input = string(b) // initialize other fields JIT as needed

	return &parser{}, nil
}

type Parser interface {
	Parse() (string, error)
}

func process(c Cases, in string) string {
	switch c {
	case none:
		return ""
	case upper: // NOW IS THE TIME FOR ALL GOOD MEN TO COME TO THE AID OF THEIR COUNTRY.
		return strings.ToUpper(in)
	case lower: // now is the time for all good men to come to the aid of their country.
		return strings.ToLower(in)
	case title: // Now Is The Time For All Good Men To Come To The Aid Of Their Country.
		return strings.ToTitle(in)
	case reverse: // nOW IS THE TIME FOR all GOOD MEN TO COME TO THE AID OF THEIR COUNTRY.
		return fold(in)
	case camel: // nowIsTheTimeForALLGoodMenToComeToTheAidOfTheirCountry.
		return spacer(in, "", true, true)
	case snake: // now_is_the_time_for_all_good_men_to_come_to_the_aid_of_their_country.
		return spacer(strings.ToLower(in), "_", true, false)
	case snakeAllCaps: // NOW_IS_THE_TIME_FOR_ALL_GOOD_MEN_TO_COME_TO_THE_AID_OF_THEIR_COUNTRY.
		return spacer(strings.ToUpper(in), "_", true, false)
	case Pascal: // NowIsTheTimeForALLGoodMenToComeToTheAidOfTheirCountry.
		return spacer(in, "", false, true)
	case kehab: // now-is-the-time-for-all-good-men-to-come-to-the-aid-of-their-country.
		return spacer(strings.ToLower(in), "-", true, false)
	case snakeCamel: // now_Is_The_Time_For_All_Good_Men_To_Come_To_The_Aid_Of_Their_Country.
		return spacer(strings.ToLower(in), "_", true, true)
	case snakePascal: // Now_Is_The_Time_For_All_Good_Men_To_Come_To_The_Aid_Of_Their_Country.
		return spacer(strings.ToLower(in), "_", false, true)
	case kehabCamel: // now-Is-The-Time-For-All-Good-Men-To-Come-To-The-Aid-Of-Their-Country
		return spacer(strings.ToLower(in), "-", true, true)
	case kehabPascal: // Now-Is-The-Time-For-All-Good-Men-To-Come-To-The-Aid-Of-Their-Country
		return spacer(strings.ToLower(in), "-", false, true)
	default:
		return in
	}
}

func NewCaseStringer(c Cases, s string) Caser {
	return &caser{s, "", c}
}

type ParserFunctionMap map[string]func()

type parser struct {
	input  string           // original input string
	lines  []string         // buffer for lines of text
	output *strings.Builder // output builder

	options ParserOptions

	dirty bool // is output modified
}

func (p *parser) Parse() (string, error) {
	return p.parse()
}

// parse contains the specific parsing algorithm
func (p *parser) parse() (string, error) {

	// var skipPrefixes []byte = []byte{'#'}

	// process 'whole text' changes
	p.setLineBreak("\n")

	// split into lines
	p.split()

	// process 'line by line' changes

	// for _, line := range bytes.Split(b, nlbyte) {
	// 	if bytes.HasPrefix(line, skipPrefixes) {
	// 		continue
	// 	}
	// }
	return p.output.String(), nil
}

func (p *parser) split() error {
	if len(p.lines) > 0 {
		return errors.New("parser.split may only be run once")
	}
	// p.lines = strings.Split(p.input, p.nl)
	return nil
}

func (p *parser) checklines() error {
	if len(p.lines) < 1 {
		err := p.split()
		if err != nil {
			return fmt.Errorf("parser could not access lines: %v", err)
		}
	}
	return nil
}

func (p *parser) setLineBreak(nl string) {
	p.options.nl = nl
}

func (p *parser) replaceLineDelimeters(old string) {
	if p.options.nl == "" {
		p.options.nl = defaultNL
	}
	p.input = strings.ReplaceAll(p.input, old, p.options.nl)

}

func (p *parser) removeSuffix() error {
	return nil
}

func (p *parser) removePrefix() error {

	for i, line := range p.lines {
		p.lines[i] = strings.TrimPrefix(line, p.options.prefix)
	}
	return nil
}

func (p *parser) skipPrefixes() error {
	return nil
}
