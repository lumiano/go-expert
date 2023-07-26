package main

import "fmt"

func sum(a, b *int) int {

	*a = 20

	return *a + *b
}

func main() {
	a := 10

	fmt.Println("a before", a)
	fmt.Println("a address", &a)

	var aPointer *int = &a

	fmt.Println("aPointer", aPointer)

	*aPointer = 20

	fmt.Println("a after", a)
	fmt.Println("aPointer", aPointer)
	fmt.Println("aPointer value", *aPointer)

	x := 10
	b := 20

	fmt.Println("sum", sum(&x, &b))

}
