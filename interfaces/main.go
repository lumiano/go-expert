package main

import "fmt"

type Account interface {
	Update() Customer
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

func (c Customer) Update() Customer {

	c.Active = !c.Active

	return c
}

func UpdateAccount(account Account) Customer {
	return account.Update()
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

	customer = UpdateAccount(customer)

	fmt.Println(customer)

}
