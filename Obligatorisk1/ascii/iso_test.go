package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestExtendedASCIIText(t *testing.T) {
	for i := 0; i<len(letters); i++ {
		if letters[i] > 0xFF {
			t.Fail()
			ting := string(letters[i])
			fmt.Printf("%s is not in the extended ASCII table\n", ting)
		} else if letters[i] < 0x80 {
			t.Fail()
			ting := string(letters[i])
			fmt.Printf("%s is not in the extended ASCII table,\n but it is in the normal ASCII table\n", ting)
		}
	}
}

func ISO(input string) {
	//splits input into a string array
	splitString := []string(strings.Split(input, ""))
	//takes each character and converts to rune, then converts to int
	for i := 0; i < len(splitString); i++ {
		character := splitString[i]
		CharToRune := []rune(character)
		RuneToInt := int(CharToRune[0])
		//prints whether its part of ASCII extended or not
		if RuneToInt >= 0x80 && RuneToInt <= 0xFF {
			fmt.Printf("%s is in extended ASCII table at position %d\n", character, RuneToInt)
		} else {
			fmt.Printf("%s is NOT in extended ASCII table (position %d)\n", character, RuneToInt)
		}
	}
}
