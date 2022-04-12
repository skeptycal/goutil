// standard command line flags based on GNU standards.
//
// POSIX
//
//
//
// Reference: https://www.gnu.org/prep/standards/html_node/Command_002dLine-Interfaces.html
package gnuflags

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	modeNormal               = 0644
	modeDir                  = 0755
	GNU_LONG_OPTION_SPEC_URL = `https://www.gnu.org/prep/standards/html_node/Option-Table.html#Option-Table`
)

var (

	// It is a good idea to follow the POSIX guidelines for the command-line options of a program. The easiest way to do this is to use getopt to parse them. Note that the GNU version of getopt will normally permit options anywhere among the arguments unless the special argument ‘--’ is used. This is not what POSIX specifies; it is a GNU extension.
	UseDoubleDash = true
)

var (
	verboseFlag = flag.Bool("verbose", false, "")
	outputFlag  = flag.String("output", "", "output file")
)

func GetOpt() {

}

// GetOptLong returns the long long option names for flags.
//
// Please define long-named options that are equivalent to the single-letter Unix-style options. We hope to make GNU more user friendly this way. This is easy to do with the GNU function getopt_long.
//
// One of the advantages of long-named options is that they can be consistent from program to program. For example, users should be able to expect the “verbose” option of any GNU program which has one, to be spelled precisely ‘--verbose’. To achieve this uniformity, look at the table of common long-option names when you choose the option names for your program (see Option Table).
//
// Options Table: https://www.gnu.org/prep/standards/html_node/Option-Table.html#Option-Table
func GetOptLong() (string, error) {
	s, err := getLongOptionSpec(GNU_LONG_OPTION_SPEC_URL)
	if err != nil {
		return "", err
	}

	return s, nil
}

func getLongOptionSpec(url string) (string, error) {
	resp, err := http.Get(url)
	if err == nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// test output:
	os.WriteFile("./longopts.html", b, modeNormal)

	return parseLongOptions(b)
}

func parseLongOptions(b []byte) (string, error) {

	return "", nil
}

func Example() {
	fmt.Println("Example import verification from goutil/manage/gnuflags")
}
