package main

import "fmt"

type MyNumber int

type Summable interface {
	~int | ~float64
}

func Compare[T comparable](a, b T) bool {
	return a == b

}

func Sum[T Summable](m map[string]T) T {

	var sum T

	for _, v := range m {
		sum += v
	}

	return sum
}

func main() {

	m := map[string]int{
		"Keven": 12425,
		"Alex":  21556,
	}

	m2 := map[string]float64{
		"Keven": 124.245,
		"Alex":  21556.25,
	}

	m3 := map[string]MyNumber{
		"Keven": 12424,
		"Alex":  255,
	}

	fmt.Println(Sum(m))
	fmt.Println(Sum(m2))
	fmt.Println(Sum(m3))
	fmt.Println(Compare(1, 2))

}
