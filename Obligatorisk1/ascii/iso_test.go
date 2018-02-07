package main

import (
	"fmt"
	"strings"
)

func main() {
	test()
}

func test() {
	fmt.Printf("%q\n", strings.Split(" € ÷ ¾ dollar ", ","))
}