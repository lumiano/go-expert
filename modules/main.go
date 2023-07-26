package main

import (
	"fmt"
	mathematics "github.com/lumiano/go-expert/mathematics"
)

func main() {
	s := mathematics.Sum(1, 2)

	fmt.Println("Result:", s)
	fmt.Println(mathematics.HelloWorld)
	fmt.Println(mathematics.IsVisible)

	calculator := mathematics.Calculator{Accumulator: 10}

	fmt.Println(calculator)
}
