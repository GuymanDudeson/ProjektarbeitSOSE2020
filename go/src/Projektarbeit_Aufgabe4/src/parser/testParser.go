package parser
//Equivalent to C++ Sourcefile
import "fmt"

func display(e Optional) {
	if e.isNothing(){
		fmt.Printf("nothing\n")
	} else {
		fmt.Printf("%s \n", e.fromJust().pretty())
	}
	return
}

func TestParserGood() []Optional {		//Parses test strings into ASTs and displays them in Terminal.
	var ast Optional
	var parsedAst []Optional			//ASTs are saved in a slice that is returned for VM test cases

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