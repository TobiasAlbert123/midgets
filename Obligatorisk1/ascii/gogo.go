package main

import (
	"fmt"
	"strings"
)

func main() {
	test()
}

func test() {
	//fmt.Println("\" € ÷ ¾ dollar \"")
	//fmt.Printf("%q\n", strings.Split("\" € ÷ ¾ dollar \"", ""))
	riktig := 0
	a := []string(strings.Split(" € ÷ ¾ dollar ", ""))
	notfound := false
	for i:=0; i < len(a) && notfound == false; i++ {
		b := a[i]
		for t := 0; t < 256; t++{
			if b == string(t){
				riktig++
				notfound = true
			}
		}
		//fmt.Printf("%s\n", a[i])
	}
	fmt.Println(riktig)
	if riktig/len(a) != 1 {
		fmt.Println("Et symbol er ikke fra ascii listen")
	}
	fmt.Println(a)
	fmt.Println(notfound)
}