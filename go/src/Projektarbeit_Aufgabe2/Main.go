package main

import (
	"fmt"
)

func main() {
	var s = "12345"
	fmt.Printf("%s reversed normally is %s\n" , s, reverse(s))
	fmt.Printf("%s reversed recursively is %s\n" , s, reverseRec(s, 0))
}


//Teil a
func reverse(s string) string{
	var reversedString = []rune("")

	for i := len(s) - 1; i >= 0; i-- {
		reversedString = append(reversedString, rune(s[i]))
	}

	return string(reversedString)
}

//Teil b

func reverseRec(s string, n int) string{
	if n != len(s) {
		return putBack(reverseRec(s, n + 1), rune(s [n]))
	} else {
		return ""
	}
}

//Util

func putBack(s string, c rune) string{
	var r = []rune(s)
	r = append(r, c)
	return string(r)
}
