package main

import (
	"fmt"

	"./accounts"
	"./clients"
)

func main() {
	johnAccount := accounts.CurrentAccount{
		Holder:  clients.Client{"v12", "John"},
		Branch:  10,
		Account: 5,
		Balance: 9,
	}
	fmt.Println(johnAccount)
	fmt.Println(PayBill(&johnAccount, 5))
	fmt.Println(johnAccount)

	jackAccount := accounts.SavingsAccount{
		Holder:  clients.Client{"v13", "Jack"},
		Branch:  101,
		Account: 52,
		Balance: 90,
	}

	fmt.Println(jackAccount)
	fmt.Println(PayBill(&jackAccount, 50))
	fmt.Println(jackAccount)
}

func PayBill(account WithdrawTransaction, value float64) string {
	return account.Withdraw(value)
}

type WithdrawTransaction interface {
	Withdraw(value float64) string
}
