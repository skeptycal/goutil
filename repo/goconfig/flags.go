package goconfig

import "flag"

var (
	versionFlag bool
	outFileFlag string
)

func init() {
	if !flag.Parsed() {

		flag.StringVar(&outFileFlag, "out", "", "name of output file")

		flag.BoolVar(&versionFlag, "version", false, "show app version and exit")

		flag.Parse()
	}
}
