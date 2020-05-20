package testParser

import (
	"Projektarbeit_Aufgabe4/src/parser"
	"Projektarbeit_Aufgabe4/src/utility"
	"fmt"
)

func Display(e utility.Optional) {
	if e.IsNothing(){
		fmt.Printf("nothing\n")
	} else {
		fmt.Printf("%s \n", e.FromJust().Pretty())
	}
	return
}

func TestParserGood() {
	Display(parser.NewParser("1").Parse())
	Display(parser.NewParser("1 + 0").Parse())
	Display(parser.NewParser("1 + (0) ").Parse())
	Display(parser.NewParser("1 + 2 * 0").Parse())
	Display(parser.NewParser("1 * 2 + 0").Parse())
	Display(parser.NewParser("(1 + 2) * 0").Parse())
	Display(parser.NewParser("(1 + 2) * 0 + 2").Parse())
}