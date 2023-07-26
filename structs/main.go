package main

import "fmt"

type HandleAccount interface {
	HandleAccount() bool
}

type Address struct {
	Street string
	Number int
}

type Customer struct {
	Name    string
	Age     int
	Active  bool
	Address Address
}

func (c Customer) HandleAccount() Customer {
	c.Active = false

	return c
}

func main() {
	customer := Customer{
		Active: true,
		Name:   "John",
		Age:    30,
		Address: Address{
			Street: "Street 2",
			Number: 0,
		},
	}

	fmt.Printf("Name: %s\n", customer.Name)
	fmt.Printf("Age: %d\n", customer.Age)
	fmt.Printf("Active: %t\n", customer.Active)

	customer.Name = "Doe"

	fmt.Println(customer)

	customer.Address.Street = "Street 1"
	customer.Address.Number = 123

	fmt.Println(customer)

	customer = customer.HandleAccount()

	fmt.Println(customer)
}
