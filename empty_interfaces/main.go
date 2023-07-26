package main

import "fmt"

func ShowType(t interface{}) {
	fmt.Printf("Type of variable is %T and value is %v\n", t, t)
}

func main() {
	var x interface{} = 10
	var y interface{} = "hello"

	ShowType(x)
	ShowType(y)

}
