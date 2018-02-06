package main

import (
	"fmt"
)

func main() {
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
