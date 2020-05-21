package ParserVM

import "fmt"

func display(e Optional) {
	if e.isNothing(){
		fmt.Printf("nothing\n")
	} else {
		fmt.Printf("%s \n", e.fromJust().Pretty())
	}
	return
}

func TestParserGood() {
	display(NewParser("1").parse())
	display(NewParser("1 + 0").parse())
	display(NewParser("1 + (0) ").parse())
	display(NewParser("1 + 2 * 0").parse())
	display(NewParser("1 * 2 + 0").parse())
	display(NewParser("(1 + 2) * 0").parse())
	display(NewParser("(1 + 2) * 0 + 2").parse())
}