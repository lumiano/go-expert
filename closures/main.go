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

	total := func() int {
		return sum(0, 6) * 50
	}()

	fmt.Println(total)
}
