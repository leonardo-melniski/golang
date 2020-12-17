package accounts

type CurrentAccount struct {
	Holder  string
	Branch  int
	Account int
	Balance float64
}

func (c *CurrentAccount) Withdrawal(value float64) string {
	if value <= 0 {
		return "Invalid input"
	}
	if c.Balance >= value {
		c.Balance -= value
		return "Success"
	}
	return "Insufficient funds"
}
