package accounts

import "../clients"

type SavingsAccount struct {
	Holder  clients.Client
	Branch  int
	Account int
	Balance float64
}

func (c *SavingsAccount) Withdraw(value float64) string {
	if value <= 0 {
		return "Invalid input"
	}
	if c.Balance >= value {
		c.Balance -= value
		return "Success"
	}
	return "Insufficient funds"
}
