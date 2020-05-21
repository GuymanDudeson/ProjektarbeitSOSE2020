package main

import (
	"ProjektarbeitSOSE2020/go/src/Projektarbeit_Aufgabe4/src/parser"
	"ProjektarbeitSOSE2020/go/src/Projektarbeit_Aufgabe4/src/vm"
)

func main() {
	//parser.TestParserGood()

	var astToTranslate []parser.Optional
	astToTranslate = append(astToTranslate, parser.NewParser("1 + 0").Parse())
	astToTranslate = append(astToTranslate, parser.NewParser("1 + 2 * 0").Parse())
	astToTranslate = append(astToTranslate, parser.NewParser("(1 + 2) * 0 + 2").Parse())
	astToTranslate = append(astToTranslate, parser.NewParser("(1 + 2) * (0 + 2) + (1 * ((0 * 2 + 2 * 1) * (2 + 0)))").Parse())
	vm.TestVM(astToTranslate)
}