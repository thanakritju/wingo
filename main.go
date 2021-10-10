package main

import (
	"flag"

	"wingo/functions"
)

var (
	function = flag.String("f", "changeBackground", "windows function to call")
)

func main() {
	flag.Parse()
	switch *function {
	case "changeBackground":
		imagePath := flag.Arg(0)
		functions.ChangeBackground(imagePath)
	default:
		println("Invalid Command")
	}
}
