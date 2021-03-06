# next step of go lang


## struct

```go
type CurrentAccount struct {
	holder  string
	branch  int
	account int
	balance float64
}

// init
account := CurrentAccount{
    holder:  "holder",
    branch:  10,
    account: 5,
    balance: 2.3,
}
account := CurrentAccount{"holder2", 100, 50, 20.3}

// with point
var account *CurrentAccount
account = new(CurrentAccount)
account.holder = "holder"

// method def and call
func (c *CurrentAccount) Withdraw(value float64) (float64, string) {
}
account.Withdraw(value)
```

## struct composition

```go
// in clients/client.go
type Client struct {
	ID   string
	Name string
}

// in accounts/currentAccounts.go
type CurrentAccount struct {
	Holder  clients.Client
	Branch  int
	Account int
	Balance float64
}
```


## interface

```go
func (c *CurrentAccount) Withdraw(value float64) (float64, string) {
}
func (c *SavingsAccount) Withdraw(value float64) (float64, string) {
}

func PayBill(account WithdrawTransaction, value float64) string {
	return account.Withdraw(value)
}

type WithdrawTransaction interface {
	Withdraw(value float64) string
}
```

`CurrentAccount` and `SavingsAccount` implements method `Withdraw` of `interface` then the function `PayBill` receive a pointer of savings accounts or current account.

```go
PayBill(&currentAccount, value)
PayBill(&savingsAccount, value)
```