package main

import (
	"fmt"
)

func main() {
	var ages [3]int
	ages[0] = 1
	ages[1] = 2
	ages[2] = 3

	fmt.Printf("%v\n", ages)

	for i, age := range ages {
		fmt.Printf("%d: %d\n", i, age)
	}
}
