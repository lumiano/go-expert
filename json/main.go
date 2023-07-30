package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Balance   float64 `json:"balance"`
}

func main() {

	account := Account{FirstName: "John", LastName: "Doe", Balance: 25616.2}

	_, err := json.Marshal(account)

	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(os.Stdout)

	err = encoder.Encode(account)
	if err != nil {
		return
	}

	var account2 Account

	f := []byte(`{"first_name":"John","last_name":"Doe","balance":22145.2}`)

	err = json.Unmarshal(f, &account2)

	if err != nil {
		return
	}

	err = encoder.Encode(account2)

	if err != nil {
		return
	}

}
