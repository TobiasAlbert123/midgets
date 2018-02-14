package main

import (
	"fmt"
	"strings"
)

func main() {
	test2()
}

/*
Inefficient method
func test() {
	//fmt.Println("\" € ÷ ¾ dollar \"")
	//fmt.Printf("%q\n", strings.Split("\" € ÷ ¾ dollar \"", ""))
	riktig := 0
	a := []string(strings.Split(" € ÷ ¾ dollar ", ""))
	notFound := false
	for i:=0; i < len(a) && notFound == false; i++ {
		b := a[i]
		for t := 0; t < 256; t++{
			if b == string(t){
				riktig++
				notFound = true
			}
		}
		//fmt.Printf("%s\n", a[i])
	}
	fmt.Println(riktig)
	if riktig/len(a) != 1 {
		fmt.Println("Et symbol er ikke fra ascii listen")
	}
	fmt.Println(a)
	fmt.Println(notFound)
}
*/

func test2() {
	splitString := []string(strings.Split("\" € ÷ ¾ dollar \"", ""))
	for i:=0; i<len(splitString); i++ {
		character := splitString[i]
		CharToRune := []rune(character)
		RuneToInt := int(CharToRune[0])
		if RuneToInt >= 0x80 && RuneToInt <= 0xFF {
			a := character
			b := RuneToInt
			fmt.Printf("%s is in extended ASCII table at position %d\n", a, b)
		}	else {
				fmt.Printf("%s is NOT in extended ASCII table (position %d)\n", character, RuneToInt)
			}
	}
}