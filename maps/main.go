package main

import "fmt"

func main() {
	salaries := map[string]int{"Lucas": 1223}

	fmt.Printf("salaries: %v\n", salaries)

	for key, value := range salaries {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}

	salaries["Lucas"] = 2424

	fmt.Printf("salaries: %v\n", salaries)

	fmt.Printf("salaries: %v\n", salaries)

	salaries["Brian"] = 1234

	fmt.Printf("salaries: %v\n", salaries)

	delete(salaries, "Lucas")

	fmt.Printf("salaries: %v\n", salaries)

	salaries = make(map[string]int)
	salaries = map[string]int{}

}
