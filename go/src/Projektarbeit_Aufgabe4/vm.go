package Projektarbeit_Aufgabe4

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

func newPush(o OpCode, i int) Code{
	return Code{
		kind: VM_PUSH,
		val:  i,
	}
}

func newPlus(o OpCode) Code{
	return Code{
		kind: VM_PLUS,
		val:  0,
	}
}

func newMult(o OpCode) Code{
	return Code{
		kind: VM_MULT,
		val:  0,
	}
}

type VM struct {
	code  []Code
	stack Stack
}

func newVM(c []Code) VM{
	return VM{
		code:  c,
		stack: nil,
	}
}

func(vm VM) clearStack() {
	vm.stack.clear()
}

func(vm VM) run() int{
	vm.clearStack()

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
		return 0
	}
	return vm.stack.getTop()
}