package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	test()
}

func test() {
	fmt.Println("\" € ÷ ¾ dollar \"")
	fmt.Printf("%q\n", strings.Split("\" € ÷ ¾ dollar \"", ""))
	a := []string(strings.Split(" € ÷ ¾ dollar ", ""))
	for i:=0; i < len(a); i++ {
		b := a[i]
		c := strconv.Atoi(b)
		fmt.Println(c)
		fmt.Printf("%s\n", a[i])
	}
	fmt.Println(a)
}