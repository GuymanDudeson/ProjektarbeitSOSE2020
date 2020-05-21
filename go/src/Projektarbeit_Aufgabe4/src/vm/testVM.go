package vm

import "fmt"

func showVMRes(r int) {
	if r == -1{
		fmt.Printf("\nVM stack (top): empty")
		return
	}
	fmt.Printf("\nVM stack (top): %d", r)
}

func TestVM(){
	var vc = []Code{NewPush(1),
		NewPush(2),
		NewPush(3),
		NewMult(),
		NewPlus()}

	var res = NewVM(vc).run()
	showVMRes(res)

	vc = []Code{NewPush(2),
		NewPush(3),
		NewPush(5),
		NewPlus(),
		NewMult()}

	res = NewVM(vc).run()
	showVMRes(res)
}