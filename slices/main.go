package main

import "fmt"

func main() {

	ages := []int{1, 2, 3, 4, 5}

	fmt.Printf("ages: %v\n", ages)

	fmt.Printf("len=%d cap=%d %v\n", len(ages), cap(ages), ages)

	fmt.Printf("len=%d cap=%d %v\n", len(ages[2:]), cap(ages[2:]), ages[2:])

	fmt.Printf("ages[1:4] = %v\n", ages[1:4])

	ages = append(ages, 6, 7, 8, 9, 10)
}
