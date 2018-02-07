package main

import (
	"fmt"
)

func main() {
	//IterateOverASCIIStringLiteral()
	ExtendedASCIIText()
}

func IterateOverASCIIStringLiteral() {
	a:= ""
	b:= "\\x"
	for i:=128; i < 256; i++ {
		h := fmt.Sprintf("%X", i)
		st := string(i)
		a = a + b + h
		fmt.Printf("%s %s %d\n", h, st, i)
	}
	fmt.Println(a)

}

func ExtendedASCIIText() {
	a := string(34)
	b := string(32)
	c := string(8364)
	d := string(247)
	e := string(190)
	f := string(100)
	g := string(111)
	h := string(108)
	i := string(97)
	j := string(114)
	zz := a+b+c+b+d+b+e+b+f+g+h+h+i+j+b+a
	fmt.Printf("%s", zz)
}