package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "12345"
	fmt.Printf("%s reversed normally is %s\n" , s, reverse(s))
	fmt.Printf("%s reversed recursively is %s\n" , s, reverseRec(s, 0))
}


//Teil a
func reverse(s string) string{
	var reversedString = strings.Builder{}

	for i := len(s) - 1; i >= 0; i-- {
		reversedString.WriteByte(s[i])
	}

	return reversedString.String()
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
	var newString strings.Builder

	newString.WriteString(s)
	newString.WriteRune(c)

	return newString.String()
}
