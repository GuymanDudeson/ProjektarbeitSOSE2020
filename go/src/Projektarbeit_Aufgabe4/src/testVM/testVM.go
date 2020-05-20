package testVM

import (
	"Projektarbeit_Aufgabe4/src/vm"
	"fmt"
)

func ShowVMRes(r int) {
	if r == -1{
		fmt.Printf("\nVM stack (top): empty")
		return
	}
	fmt.Printf("\nVM stack (top): ", r)
}

func TestVM(){
	var vc = []vm.Code{vm.NewPush(1),
					vm.NewPush(2),
					vm.NewPush(3),
					vm.NewMult(),
					vm.NewPlus()}

	var res = vm.NewVM(vc).Run()
	ShowVMRes(res)

	vc = []vm.Code{vm.NewPush(2),
		vm.NewPush(3),
		vm.NewPush(5),
		vm.NewPlus(),
		vm.NewMult()}

	res = vm.NewVM(vc).Run()
	ShowVMRes(res)
}