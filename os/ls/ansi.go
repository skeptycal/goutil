package main

import "fmt"

const ansiFMT string = `\u001b[%dm`

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed    Color = "\u001b[31m"
	ColorGreen  Color = "\u001b[32m"
	ColorYellow Color = "\u001b[33m"
	ColorBlue   Color = "\u001b[34m"
	ColorReset  Color = "\u001b[0m"
)

func (c Color) String() string {
	return string(c)
}

func ansiFmt(i int) string {
	return fmt.Sprintf(ansiFMT, i)
}

func CPrint(color Color, message string) {
	fmt.Println(color, message, ColorReset)
}
