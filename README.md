# ProjektarbeitSOSE2020

### This Repository is a project for University.
### The Assignment is to implement various functions/programs initially meant to teach and better understand C/C++ in GO.

**Assignment 0 Prologue:**
- Added to explain the significance of rune as a datatype for string operations in rune.
- Important to understand design decisions in Assignment 1 and 2.
 
**Assignment 1:**
- Finding and extracting the longest suffix of a String after a predefined Stringpattern.
- This is supposed to show pointer arithmetic and string handling in C.
- GO has no pointer arithmethic and strings are represented as a slice of bytes.
- Changes made were:
  - Representing a string as a slice of runes to prevent issues when handling Non-ASCII characters
  - New logic without pointer arithmetic.
- GO also provides a function library for strings which includes a SplitAfter() function:
  - This function creates a new slice out of substrings separated by a pattern.
  - The last entry in the slice is automatically what we are searching for; the longest suffix.

**Assignment 2:**
- Reversing a string iteratively and recursively
- This is supposed to reinforce pointer arithmetic and stringbuilding in general. Also a refresher on recursive functions.
- GO has no pointer arithmetic and stringbuilding is simple (no nullterminator and no memory allocation if reslicing is possible)
- Changes made were:
  - Representing a string as a slice of runes to handle Non-ASCII characters.
  - New logic without pointer arithmetic

**Assigment 3:**
  - Too C++ specific. No sensible way to translate to GO, therefore discarded.

**Assigment 4:**
- Building a parser for simple arithmetic strings which builds a mathematically correct tree for the order of operations.
- Building a VM which can interpret a parsed tree and calculate the result.
- Parser and VM are provided in the original assigment.
- This is supposed to show how you import projects, handle makefiles and parser/VM functionality.
- GO doesn't use Makefiles but Directories for it's dependencies. Stacks are not a standard structure type. 
  Enums and inheritance are structured differently. Doesn't support generics.
- Changes made were:
  - Translated all C-sourcecode to GO where possible.
  - Stack realised as a custom struct using a slice. Appending = Push, Removing last index = Pop.
  - Enums realised as custom types whith constants assigned to them.		
  - Inheritance realised with interfaces.
  - Vectors realised as slices.
  - Changed generics to static types.