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
```