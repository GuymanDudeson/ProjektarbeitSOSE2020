package main

import (
	"fmt"
)

func main() {
	var s = "12345"
	fmt.Printf("%s reversed normally is %s\n" , s, reverse(s))
	fmt.Printf("%s reversed recursively is %s\n" , s, reverseRec(s))

	s = "ΖΠӲפ"
	fmt.Printf("%s reversed normally is %s\n" , s, reverse(s))
	fmt.Printf("%s reversed recursively is %s\n" , s, reverseRec(s))

	s = "Ζ2Ӳa"
	fmt.Printf("%s reversed normally is %s\n" , s, reverse(s))
	fmt.Printf("%s reversed recursively is %s\n" , s, reverseRec(s))
}


//Teil a
//Runs through the string from back to front
//Appends each byte to the end of return string
func reverse(s string) string{
	var originalString = []rune(s)
	var reversedString []rune

	for i := len(originalString) - 1; i >= 0; i-- {
		reversedString = append(reversedString, originalString[i])
	}

	return string(reversedString)
}

//Teil b
//Recursively appends the first Character of the given String to it's end.
//Continues Recursion with the Original String without the first Character
//Recursion ends when no Characters are left in String (length == 0)
func reverseRec(s string) string{
	var originalString = []rune(s)
	if len(originalString) == 0 {
		return ""
	} else {
		return putBack(reverseRec(string(originalString[1:])), originalString[0])
	}
}

//Util

func putBack(s string, c rune) string{
	var r = []rune(s)
	r = append(r, c)
	return string(r)
}
