package stack

type Stack struct {
	slice []int
}

func(stack Stack) Clear(){
	stack.slice = stack.slice[:0]
}

func(stack Stack) IsEmpty() bool {
	return len(stack.slice) == 0
}

func(stack Stack) GetTop() int{
	return stack.slice[len(stack.slice) - 1]
}

func(stack Stack) Push(val int) {
	stack.slice = append(stack.slice, val)
}

func(stack Stack) Pop() {
	stack.slice = stack.slice[:len(stack.slice)-1]
}