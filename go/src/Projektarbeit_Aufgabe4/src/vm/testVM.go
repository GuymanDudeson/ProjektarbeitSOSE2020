package vm

import (
	"ProjektarbeitSOSE2020/go/src/Projektarbeit_Aufgabe4/src/parser"
	"fmt"
)

func showVMRes(r int) {
	if r == -1{
		fmt.Printf("\nVM stack (top): empty")
		return
	}
	fmt.Printf("\nVM stack (top): %d", r)
}

func TestVM(astToTranslate []parser.Optional){
	var astCode []Code
	var res int
	for _, att := range astToTranslate{			//Takes the already parsed AST builds VM code out of it and runs a new VM
		astCode = AstToCode(att, astCode)
		res = NewVM(astCode).run()
		showVMRes(res)
		astCode = astCode[:len(astCode) - 1]
	}

	var vc = []Code{NewPush(1),
		NewPush(2),
		NewPush(3),
		NewMult(),
		NewPlus()}

	res = NewVM(vc).run()
	showVMRes(res)

	vc = []Code{NewPush(2),
		NewPush(3),
		NewPush(5),
		NewPlus(),
		NewMult()}

	res = NewVM(vc).run()
	showVMRes(res)
}