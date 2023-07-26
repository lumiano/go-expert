package main

import "fmt"

func sum(numbers ...int) int {
	s := 0

	for _, n := range numbers {
		s += n
	}

	return s
}

func main() {

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
