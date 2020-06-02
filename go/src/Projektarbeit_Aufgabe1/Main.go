package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "Ha::ll:o"
	var p = "::"


	fmt.Printf("The longest Suffix of %s that does not contain %s is %s\n", s, p, longestSuffixEasy(s, p))
	fmt.Printf("The longest Suffix of %s that does not contain %s is %s\n", s, p, longestSuffixWithoutSplit(s, p))

	s = "Martin Sulzmann"
	p = " "

	fmt.Printf("The Surname of %s is %s\n", s, longestSuffixEasy(s, p))
	fmt.Printf("The Surname of %s is %s\n", s, longestSuffixWithoutSplit(s, p))

	s = "Ottottottottotto"
	p = "otto"

	fmt.Printf("The longest Suffix of %s that does not contain %s is %s\n", s, p, longestSuffixEasy(s, p))
	fmt.Printf("The longest Suffix of %s that does not contain %s is %s\n", s, p, longestSuffixWithoutSplit(s, p))

	var tests = []TestCase{
		{longestSuffixWithoutSplit("Ha::ll:o", "::"), ":ll:o"},
		{longestSuffixWithoutSplit("Ha::ll::o", "::"), ":o"},
		{longestSuffixWithoutSplit("Ha:ll:o", "::"), "Ha:ll:o"},
		{longestSuffixWithoutSplit("Ha::ll:::o", "::"), ":o"},
		{longestSuffixWithoutSplit("Ha::ll:o::", "::"), ":"},
		{longestSuffixWithoutSplit("Haϴϴll:oϴϴ", "ϴϴ"), "ϴ"},
		{longestSuffixWithoutSplit("HaϴϴΈlΈo", "ϴϴ"), "ϴΈlΈo"},
		{longestSuffixWithoutSplit("ϴΈ::Έ", "::"), ":Έ"},
	}
	runTests(tests)
}

func longestSuffixEasy(s string, pattern string) string {
	var suffixes = strings.SplitAfter(s, pattern) //Creates a new Slice with substrings of s separated by the specified pattern
	return suffixes[len(suffixes) - 1]		//Last Slice entry is automatically the longest suffix without the pattern
}

//2 runs over the String as a []rune
//1. Search possible Index for the longest suffix
//2. Check if Index represents entire pattern
//return resliced string only containing longest suffix
func longestSuffixWithoutSplit(s string, pattern string) string{
	var originalString = []rune(s)
	var stringLength = len(originalString)

	var patternRunes = []rune(pattern)
	var patternlength = len(patternRunes)

	var longestSuffixIndex = 0

	//Loop iterates over all runes in the give String
	for i, r := range originalString {

		//If the remaining String is not big enough to contain the Pattern stop searching
		if stringLength - i >= patternlength{

			//If a rune matches the start of the pattern compare all subsequent runes until one doesn't match or pattern complete
			if r == patternRunes[0] {
				for j := 1; j <= patternlength; j++ {
					if j == patternlength{
						longestSuffixIndex = i + 1
						break
					} else if originalString[i + j] != patternRunes[j]{
						break
					}
				}
			}
		}
	}

	longestSuffix := originalString[longestSuffixIndex:]

	return string(longestSuffix)
}


//Tests
type Result int
const (
	OK   Result = 1
	FAIL Result = 0
)

type TestCase struct {
	input string
	expected string
}

func runTests(tests []TestCase) {
	var result Result
	for i, t := range tests {
		fmt.Printf("Test %d\n", i)
		result = stringIsEqual(t.input, t.expected)

		if OK == result{
			fmt.Printf("OK \n")
		}
		if FAIL == result{
			fmt.Printf("FAIL \n")
		}
	}
}

func stringIsEqual(input string, expected string) Result {
	var r Result

	if strings.Compare(input, expected) == 0{
		r = OK
	} else {
		r = FAIL
	}

	return r
}



