package vm

import "Projektarbeit_Aufgabe4/src/stack"

type OpCode int

const (
	VM_PUSH OpCode = 0
	VM_PLUS OpCode = 1
	VM_MULT OpCode = 2
)

type Code struct {
	kind OpCode
	val int
}

func NewPush(i int) Code{
	return Code{
		kind: VM_PUSH,
		val:  i,
	}
}

func NewPlus() Code{
	return Code{
		kind: VM_PLUS,
		val:  0,
	}
}

func NewMult() Code{
	return Code{
		kind: VM_MULT,
		val:  0,
	}
}

type VM struct {
	code  []Code
	stack stack.Stack
}

func NewVM(c []Code) VM{
	return VM{
		code:  c,
		stack: stack.Stack{},
	}
}

func(vm VM) ClearStack() {
	vm.stack.Clear()
}

func(vm VM) Run() int{
	vm.ClearStack()

	for	_, c := range vm.code{
		switch c.kind {
		case VM_PUSH:
			vm.stack.Push(c.val)
			break
		case VM_MULT:
			var right = vm.stack.GetTop()
			vm.stack.Pop()
			var left = vm.stack.GetTop()
			vm.stack.Pop()
			vm.stack.Push(left * right)
			break
		case VM_PLUS:
			var right = vm.stack.GetTop()
			vm.stack.Pop()
			var left = vm.stack.GetTop()
			vm.stack.Pop()
			vm.stack.Push(left + right)
			break
		}
	}

	if vm.stack.IsEmpty() {
		return -1
	}
	return vm.stack.GetTop()
}