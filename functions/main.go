package main

import (
	"errors"
	"fmt"
)

func sum(a, b int) (int, error) {

	s := a + b

	if s > 10 {
		return 0, errors.New("sum is greater than 10")
	}

	return s, nil
}

func main() {
	v, e := sum(0, 6)

	if e != nil {
		fmt.Println(e.Error())
	}

	fmt.Println(v)
}
