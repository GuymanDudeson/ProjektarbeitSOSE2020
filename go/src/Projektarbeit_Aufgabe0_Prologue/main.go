package main

import "fmt"

/*
This functions as a prologue for the actual assignments that are part of this work.
The motivation for this prologue is the introduction of a new type called "rune" in GO.
This type is used when handling strings to distinguish between separate bytes and Unicode Codesamples.
Since assignments 1 and 2 work with strings and more importantly loops over string "characters",
we need to clarify why and when runes are important.
Since runes are mainly used because of the existence of Unicode, a basic understanding of Unicode
and it's encoding methods is helpful and somewhat required.

I will provide a quick summary of the most important characteristics for this prologue.
Unicode is an IT-standard used for encoding Characters into unique Codepoints.
Unicode's goal is to map every existing character in every language, compared to only english characters mapped in ASCII.
Because of the sheer amount of languages and therefor characters, one byte is no longer big enough to represent everyone of them.
That means characters in Unicode can use up to 4 bytes of memory.
Important to know is that the old ASCII characters keep their byte values and can therefore all still be represented by one byte, even in Unicode.

In the following I will show some examples of how this leads to problems when looping over strings.
 */

func main() {
	/*
	First lets compare how ASCII and Unicode Characters are represented when we output them with different formats.
	Comparing the latin character a to the Cherokee character Ꮡ
	From top to bottom we see as output:
	-The uninterpreted bytevalue of the string (byteslice)
	-the bytecode representation base 2
	-The Hexvalue base 16
	-The Unicode codepoint
	-the character literal
	 */
	ascii := 'a'
	fmt.Printf("Character a: \n")
	fmt.Printf("bytevalue: %d\n", ascii)
	fmt.Printf("base 2:    %b\n", ascii)
	fmt.Printf("base 16:   %X\n", ascii)
	fmt.Printf("unicode:   %U\n", ascii)
	fmt.Printf("literal:   %+q\n", ascii)


	fmt.Printf("\n")


	unicode := 'Ꮡ'
	fmt.Printf("Character Ꮡ: \n")
	fmt.Printf("bytevalue: %d\n", unicode)
	fmt.Printf("base 2:    %b\n", unicode)
	fmt.Printf("base 16:   %X\n", unicode)
	fmt.Printf("unicode:   %U\n", unicode)
	fmt.Printf("literal:   %q\n", unicode)


	fmt.Printf("\n")


	/*
	What we can see is, that a has the bytevalue of 97, which has the Hexvalue 61, which means 1 byte is enough to represent it.
	Ꮡ on the other hand has the bytevalue of 5073 and is represented by 3 Hexvalue-Bytes: E1 8F 91 as seen below
	 */
	fmt.Printf("Ꮡ has hexvalues:")
	for i := 0; i < len("Ꮡ");i++ {
		fmt.Printf("% X", "Ꮡ"[i])
	}


	fmt.Printf("\n\n")


	/*
	These character literals are automatically saved as a int32(spoilers: a rune) which can accurately represent any Unicode codepoint
	so we have no issue outputting them.
	But now for strings:

	Strings in GO are not implemented as a Character-pointer in c but as a slice of bytes.
	Which means we can safely cast from string to byte-slice and back.
	*/

	asciiString := "Hello World"
	asciiByteSlice := []byte(asciiString)
	fmt.Printf("Pure ASCII: %s\n", string(asciiByteSlice))
	/*
	This also works with Non-ASCII
	 */
	NonAsciiString := "Hello Ꮡ-Ꮡ"
	NonAsciiByteSlice := []byte(NonAsciiString)
	fmt.Printf("Non-ASCII: %s\n", string(NonAsciiByteSlice))


	fmt.Printf("\n")


	/*
	It get's interesting once we start working with the byte-slice
	Let's remove the last entry in the slice. We would expect to get the String "Hello Ꮡ-"
	 */
	foo := "Hello Ꮡ-Ꮡ"
	foo = foo[:len(foo) - 1]
	fmt.Printf("Last byte entry removed: %s\n", foo)

	fmt.Printf("\n")
	/*
	But the real output is: Hello Ꮡ-�� (Or other end characters depending on environment)
	Not what we expected.
	Let's look at the length of the foo-slice to find out what's happening. Length should be: 9 Right?
	*/
	foo = "Hello Ꮡ-Ꮡ"
	fmt.Printf("Lenght of foo as byte-slice: %d\n", len(foo))
	/*
	In reality it's 13.
	The reason for that is how the byte slice works for Non-ASCII Characters.
	If we remember the start of this prologue where we looked at the characteristics of Ꮡ,
	we will see that Ꮡ needed 3 bytes of memory to represent the 3 Hexvalues E1 8F 91.
	That means the Byte slice has 4 additional bytes for the 2 Non-ASCII characters.
	You can probably see why this leads to issues when using a loop to iterate over the string.
	You would have to account for characters which use different amounts of array-entries and rebuild them accordingly
	whenever you change something.
	Which brings us to Runes and why GO needs them.
	 */
	var bar = []rune("Hello Ꮡ-Ꮡ")
	/*
	A rune is a Int32 value that provides additional functions to handle Unicode Codepoints.
	A rune-slice, like initialized above, for example automatically identifies multiple bytes belonging to one character
	and safes them as a single rune(int32).
	Let's run our length function again and see if that changes anything
	 */
	fmt.Printf("Lenght of foo as rune-slice: %d\n", len(bar))
	/*
	As expected it now returns 9 as there are 9 int32 values in the slice.
	Which means if we now remove the last index in our slice we should also get our wanted: Hello Ꮡ-
	Of course, before we can output it we have to cast it back to a string.
	 */
	bar = bar[:len(bar) - 1]
	fmt.Printf("Last rune entry removed: %s\n", string(bar))

	fmt.Printf("\n")


	/*
	It is worth mentioning that GO sometimes automatically casts to runes to make certain operations possible.
	For example GOs for-each loop equivalent "for-range" will automatically iterate over a string as a rune-slice
	to prevent issues with Non-ASCII characters.
	 */
	fmt.Printf("For-each iteration over string: \n")
	forEachString := "Ꮡ-Ꮡ Hello Ꮡ-Ꮡ"
	for _, r := range forEachString{
		fmt.Printf("%s\n", string(r))
	}

	/*
	In conclusion:
	GO's representation of Strings as byte-slices causes issues when working with Non-ASCII characters.
	Rune was implemented to solve these issues and should be used whenever it is not certain if all used
	characters in a program will be ASCII.
	This is almost always the case when working in an international environment.
	 */
}