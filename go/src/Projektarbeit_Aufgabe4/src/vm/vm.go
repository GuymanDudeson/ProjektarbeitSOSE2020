package vm

import (
	"ProjektarbeitSOSE2020/go/src/Projektarbeit_Aufgabe4/src/parser"
)

type OpCode int

const (
	VM_PUSH OpCode = 0
	VM_PLUS OpCode = 1
	VM_MULT OpCode = 2
)

type Code struct {
	kind OpCode
	val  int
}

func NewPush(i int) Code {
	return Code{
		kind: VM_PUSH,
		val:  i,
	}
}

func NewPlus() Code {
	return Code{
		kind: VM_PLUS,
		val:  0,
	}
}

func NewMult() Code {
	return Code{
		kind: VM_MULT,
		val:  0,
	}
}

//Assignment Semantik

func AstToCode(astOptional parser.Optional, code []Code) []Code{	//functions as a startpoint for recursively translating the AST
	return translateAstToCode(astOptional.Val, &code)
}

func translateAstToCode(ast parser.Exp, code *[]Code) []Code {		//Works recursively through the AST adding PUSH,PLUS,MULT respectively to a slice
	if ast.GetKind() == 0 {									//We know current Exp is an IntExp and is therefore a leaf of our tree
		return append(*code, NewPush(ast.Eval()))			//so we create a new Push with its value
	} else {
		*code = translateAstToCode(ast.GetE2(), code)		//We know current Exp is a Plus/Mult and therefore a node of our tree
		*code = translateAstToCode(ast.GetE1(), code)		//so we start the recursion for its subExp
		if ast.GetKind() == 1 {		//We know current node is a Plus so we create Plus code
			return append(*code, NewPlus())
		}
		if ast.GetKind() == 2 {		//We know current node is a Plus so we create Mult code
			return append(*code, NewMult())
		}
	}
	return *code
}

//
type VM struct {
	code  []Code
	stack Stack
}

func NewVM(c []Code) *VM {
	return &VM{
		code:  c,
		stack: Stack{},
	}
}

func(vm *VM) clearStack() {  	//Empties the current stack
	vm.stack.clear()
}

func(vm *VM) run() int{
	vm.clearStack()		//Empties the current stack on a fresh start

	for	_, c := range vm.code{
		switch c.kind {
		case VM_PUSH:
			vm.stack.push(c.val)
			break
		case VM_MULT:
			var right = vm.stack.getTop()
			vm.stack.pop()
			var left = vm.stack.getTop()
			vm.stack.pop()
			vm.stack.push(left * right)
			break
		case VM_PLUS:
			var right = vm.stack.getTop()
			vm.stack.pop()
			var left = vm.stack.getTop()
			vm.stack.pop()
			vm.stack.push(left + right)
			break
		}
	}

	if vm.stack.isEmpty() {
		return -1
	}
	return vm.stack.getTop()
}