package String

import (
	"fmt"
)

type MyString struct {
	size       int
	characters []rune
}

type MyStringHandler interface {
	NewMyString(s string) *MyString
	NewMyString2(s MyString) *MyString
	PrintString() int
	Concat(s MyString) bool
	Equals(s MyString) bool
	Overwrite(s MyString) bool
	GetChar(i int) rune
	GetSize() int
}

//Creates a new MyString, allocates a new Slice for it's Runes, and fills slice with given string
func NewMyString(s string) *MyString {
	var newString = new(MyString)
	newString.characters = make([]rune, len(s))

	for i, r := range s {
		newString.characters[i] = r
	}
	newString.size = len(s)
	return newString
}

//Creates a new MyString and assigns the given Values to it
func NewMyString2(s MyString) *MyString {
	var newString = new(MyString)
	newString.characters = s.characters
	newString.size = s.size
	return newString
}

//Prints every Rune in the character Slice of the MyString,
//It returns the number of printed runes
func (s *MyString) PrintString() int {
	var no = 0

	for i, r := range s.characters {
		fmt.Printf("%s", string(r))
		no = i
	}
	fmt.Printf("\n")

	return no
}

//Returns the size of the MyString
func (s *MyString) GetSize() int {
	return s.size
}

//Returns the Rune at the index i
func (s *MyString) GetChar(i int) rune {
	return s.characters[i]
}

//Concatenates the toConcat string to the end of the MyString
//Allocates a bigger rune slice for MyString if cap is not sufficient
func (s *MyString) Concat(toConcat MyString) bool {
	s.appendCharacters(toConcat.characters)
	s.size += toConcat.size
	return true
}

//Subsequently compares every rune in the rune slices for equality,
//It returns true if both are the same and false if they are not
func (s *MyString) Equals(toCompare MyString) bool {
	if s.size != toCompare.size {
		return false
	} else {
		for i, r := range s.characters {
			if r != rune(toCompare.characters[i]) {
				return false
			}
		}
	}
	return true
}

//Overwrites MyString rune slice and size with given params
//It returns true if successful
func (s *MyString) Overwrite(newString MyString) bool {
	var newStringArray []rune
	s.characters = newStringArray
	s.size = 0
	s.appendCharacters(newString.characters)
	s.size = newString.size
	return true
}

//Appends given runes to end of MyString rune slice
//Allocates a bigger rune slice if necessary
func (s *MyString) appendCharacters(toAppend []rune) {
	var oldChars = s.characters
	var appendingChars = toAppend

	var newCharSlice []rune

	if cap(s.characters) < s.size+len(appendingChars) {
		newCharSlice = make([]rune, s.size+len(appendingChars))
	} else {
		newCharSlice = s.characters
	}

	for i := 0; i < s.size; i++ {
		newCharSlice[i] = oldChars[i]
	}

	for i := 0; i < len(appendingChars); i++ {
		newCharSlice[i+s.size] = toAppend[i]
	}

	s.characters = newCharSlice
}
