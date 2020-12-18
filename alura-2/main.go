package main

import (
	"fmt"

	"./accounts"
	"./clients"
)

func main() {
	account := accounts.CurrentAccount{
		Holder:  clients.Client{"v12", "name"},
		Branch:  10,
		Account: 5,
		Balance: 9,
	}
	fmt.Println(account)
	fmt.Println(account.Withdraw(1.1))
	fmt.Println(account)
}
