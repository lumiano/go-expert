package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	numbers := []int{1, 2, 3, 4, 5}

	for i, v := range numbers {
		fmt.Println("index:", i, "value:", v)
	}

	x := 0

	for x < 10 {
		fmt.Println(x)
		x++
	}

	for {
		fmt.Println("loop")
		break
	}
}
