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

func TestParserGood() []Optional {
	var ast Optional
	var parsedAst []Optional

	ast = NewParser("1").Parse()
	parsedAst = append(parsedAst, ast)
	display(ast)

	ast = NewParser("1 + 0").Parse()
	parsedAst = append(parsedAst, ast)
	display(ast)

	ast = NewParser("1 + (0) ").Parse()
	parsedAst = append(parsedAst, ast)
	display(ast)

	ast = NewParser("1 + 2 * 0").Parse()
	parsedAst = append(parsedAst, ast)
	display(ast)

	ast = NewParser("1 * 2 + 0").Parse()
	parsedAst = append(parsedAst, ast)
	display(ast)

	ast = NewParser("(1 + 2) * 0").Parse()
	parsedAst = append(parsedAst, ast)
	display(ast)

	ast = NewParser("(1 + 2) * 0 + 2").Parse()
	parsedAst = append(parsedAst, ast)
	display(ast)

	ast = NewParser("(1 + 2) * (0 + 2) + (1 * ((0 * 2 + 2 * 1) * (2 + 0)))").Parse()
	parsedAst = append(parsedAst, ast)
	display(ast)

	return parsedAst
}