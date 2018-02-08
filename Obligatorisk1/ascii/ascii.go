package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i:=128; i < 256; i++ {
		//IterateOverASCIIStringLiteral(i)
	}
	//fmt.Println(ExtendedASCIIText())
	letters := [16]int{34,32,8364,247,32,190,32,100,111,108,108,97,114,32,34}
	ExtendedASCIIText(letters)
}

func IterateOverASCIIStringLiteral(i int) {
	hex := fmt.Sprintf("%X", i)
	st := string(i)
	n := int64(i)
	bin := strconv.FormatInt(n, 2)
	fmt.Printf("%s %s %s\n", hex, st, bin)
}
func ExtendedASCIIText(letters [16]int) {
	allLetters := string("")
	for i:= 0; i < len(letters); i++ {
		allLetters += string(letters[i])
	}
	fmt.Printf("%s", allLetters)
}