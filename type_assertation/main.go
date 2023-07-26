package main

import "fmt"

func main() {
	var v interface{} = 10

	res, ok := v.(int)

	if ok {
		fmt.Println("res", res)
	}

}
