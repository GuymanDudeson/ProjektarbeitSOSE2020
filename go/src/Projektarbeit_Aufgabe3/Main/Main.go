package main

import (
	"Projektarbeit_Aufgabe3/String"
	"fmt"
)

func main() {
	var s1 = String.NewMyString("")
	var s2 = String.NewMyString("Hello")
	var s3 = String.NewMyString2(*s2)
	s1.Concat(*s2)
	s2.Overwrite(*s3)

	fmt.Printf("Expected Result is: Hello\n")
	fmt.Printf("Actual Result: ")
	s2.PrintString()

	fmt.Printf("\n")

	fmt.Printf("Expected Result is: l\n")
	fmt.Printf("Actual Result: ")
	fmt.Printf("%s\n", string(s2.GetChar(2)))

	fmt.Printf("\n")

	fmt.Printf("Expected Result is: Hello\n")
	fmt.Printf("Actual Result: ")
	s1.PrintString()

	fmt.Printf("\n")

	var s4 = String.NewMyString("World")
	s2.Concat(*s4)

	fmt.Printf("Expected Result is: HelloWorld\n")
	fmt.Printf("Actual Result: ")
	s2.PrintString()

	fmt.Printf("\n")

	s2.Overwrite(*s4)

	fmt.Printf("Expected Result is: World\n")
	fmt.Printf("Actual Result: ")
	s2.PrintString()

	fmt.Printf("\n")

	fmt.Printf("Expected Result is: true\n")
	fmt.Printf("Actual Result: ")
	fmt.Printf("%t\n", s3.Equals(*s1))

	fmt.Printf("\n")

	fmt.Printf("Expected Result is: false\n")
	fmt.Printf("Actual Result: ")
	fmt.Printf("%t\n", s2.Equals(*s1))
}
