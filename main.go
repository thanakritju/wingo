package main

import (
	"flag"
	"strings"

	"wingo/functions"
)

var (
	function = flag.String("f", "changebackground", "windows function to call")
)

func main() {
	flag.Parse()
	switch strings.ToLower(*function) {
	case "changebackground":
		imagePath := flag.Arg(0)
		functions.ChangeBackground(imagePath)
	case "messagebox":
		title := flag.Arg(0)
		message := flag.Arg(1)
		functions.MessageBox(title, message)
	default:
		println("Invalid Command")
	}
}
