package main

import "fmt"

type IAccount interface {
	Deposit(amount int)
	Add() *Account
}

type Account struct {
	owner   string
	balance int
}

func (a *Account) Add() *Account {

	return &Account{owner: a.owner, balance: 0}
}

func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func Deposit(account IAccount, amount int) IAccount {
	account.Deposit(amount)

	return account
}

func Add(account IAccount) IAccount {
	return account.Add()
}

func main() {

	account := &Account{
		owner: "Mike",
	}

	addAccount := Add(account)

	fmt.Println("addAccount", addAccount)

	deposit := Deposit(account, 100)

	fmt.Println("deposit", deposit)

}
