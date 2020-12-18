package accounts

import "../clients"

type CurrentAccount struct {
	Holder  clients.Client
	Branch  int
	Account int
	Balance float64
}

func (c *CurrentAccount) Withdraw(value float64) string {
	if value <= 0 {
		return "Invalid input"
	}
	if c.Balance >= value {
		c.Balance -= value
		return "Success"
	}
	return "Insufficient funds"
}
