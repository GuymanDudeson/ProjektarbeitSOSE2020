package vm
//There is no native stack implementation in Go. It can be imported but for this simple assignment I wrote my own functionality.
type Stack struct {
	slice []int		//stack is represented as a Slice of integers
}

func(stack *Stack) clear(){
	stack.slice = stack.slice[:0]		//Empty's the stack
}

func(stack *Stack) isEmpty() bool {
	return len(stack.slice) == 0
}

func(stack Stack) getTop() int{
	return stack.slice[len(stack.slice) - 1]	//returns last entry in Slice which represents the top of the stack
}

func(stack *Stack) push(val int) {
	stack.slice = append(stack.slice, val)		//appending an int to the end of the slice (the top) functions as a push
}

func(stack *Stack) pop() {
	stack.slice = stack.slice[:len(stack.slice)-1]	//removing the last entry of the slice(the top) functions as a pop
}