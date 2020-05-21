package parser

import "fmt"

func display(e Optional) {
	if e.isNothing(){
		fmt.Printf("nothing\n")
	} else {
		fmt.Printf("%s \n", e.fromJust().pretty())
	}
	return
}

func TestParserGood() {
	display(NewParser("1").Parse())
	display(NewParser("1 + 0").Parse())
	display(NewParser("1 + (0) ").Parse())
	display(NewParser("1 + 2 * 0").Parse())
	display(NewParser("1 * 2 + 0").Parse())
	display(NewParser("(1 + 2) * 0").Parse())
	display(NewParser("(1 + 2) * 0 + 2").Parse())
	display(NewParser("(1 + 2) * (0 + 2) + (1 * ((0 * 2 + 2 * 1) * (2 + 0)))").Parse())
}