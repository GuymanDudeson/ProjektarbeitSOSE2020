package Projektarbeit_Aufgabe4

import "fmt"

func display(e Optional) {
	if e.isNothing(){
		fmt.Printf("nothing\n")
	} else {
		fmt.Printf("%s \n", e.fromJust().pretty())
	}
	return
}

func testParserGood() {
	display(newParser("1").parse())
	display(newParser("1 + 0").parse())
	display(newParser("1 + (0) ").parse())
	display(newParser("1 + 2 * 0").parse())
	display(newParser("1 * 2 + 0").parse())
	display(newParser("(1 + 2) * 0").parse())
	display(newParser("(1 + 2) * 0 + 2").parse())
}