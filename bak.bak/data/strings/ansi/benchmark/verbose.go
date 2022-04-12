package benchmark

import "fmt"

func verbosef(format string, args ...interface{}) {
	if defaultConfig.verbose {
		fmt.Fprintf(defaultConfig.output, format, args...)
	}
}
