package main

import (
	. "github.com/logrusorgru/aurora"
	"github.com/skeptycal/goutil/repo/defaults"
)

var dbecho = defaults.DbEcho

func main() {

	dw := defaults.DebugWriter

	// defaults.DbEcho(dbEchoPrefix + "color dbecho" + reset)
	// fmt.Printf("dbEchoPrefix: %q - %vtests%v\n", dbEchoPrefix, dbEchoPrefix, reset)

	defaults.DebugWriter.Println("dbecho fake")

	dbecho("This is a DbEcho example.")
	dw.Print(BrightBlue("This is a DbPrint example."))
	dw.Println("")
	dw.Println(BrightGreen("This is a DbPrintln example."))
	dbecho("This is a DbEcho example.")

	dw.Printf(BrightYellow("This is a DbPrintf example. (wrap--> %v <--)").String(), "the_example")
}
