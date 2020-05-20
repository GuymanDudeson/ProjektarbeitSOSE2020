package Projektarbeit_Aufgabe4

import "fmt"

func showVMRes(r int) {
	if r == -1{
		fmt.Printf("\nVM stack (top): empty")
		return
	}
	fmt.Printf("\nVM stack (top): ", r)
}

func testVM(){
	var vc = []Code{newPush(1),
					newPush(2),
					newPush(3),
					newMult(),
					newPlus()}

	var res = newVM(vc).run()
	showVMRes(res)

	vc = []Code{newPush(2),
		newPush(3),
		newPush(5),
		newPlus(),
		newMult()}

	res = newVM(vc).run()
	showVMRes(res)
}