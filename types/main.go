package main

import "fmt"

type ID int

var id ID = 1

func main() {
	println(id)
	fmt.Printf("%T\n", id)
	fmt.Printf("%v\n", id)
}
