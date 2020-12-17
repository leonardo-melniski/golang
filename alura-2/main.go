package main

import (
	"fmt"

	"./accounts"
)

func main() {
	account := accounts.CurrentAccount{
		Holder:  "holder",
		Branch:  10,
		Account: 5,
		Balance: 9,
	}
	fmt.Println(account)
	fmt.Println(account.Withdrawal(1.1))
	fmt.Println(account)
}
