package ParserVM

type Stack struct {
	slice []int
}

func(stack Stack) clear(){
	stack.slice = stack.slice[:0]
}

func(stack Stack) isEmpty() bool {
	return len(stack.slice) == 0
}

func(stack Stack) getTop() int{
	return stack.slice[len(stack.slice) - 1]
}

func(stack Stack) push(val int) {
	stack.slice = append(stack.slice, val)
}

func(stack Stack) pop() {
	stack.slice = stack.slice[:len(stack.slice)-1]
}