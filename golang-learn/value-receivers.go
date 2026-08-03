package main

import (
	"fmt"
)

type Account struct {
	Owner  string
	Amount float64
}

// value receivers
func (a Account) DisplayBalance() {
	fmt.Printf("%s has $%.2f\n", a.Owner, a.Amount)
}

// pounter receivers
func (a *Account) Deposit(amount float64) {
	a.Amount += amount
}

func main() {
	acc := Account{Owner: "Burhan", Amount: 100.0}
	acc.DisplayBalance()
	acc.Deposit(50.0)
	acc.DisplayBalance()
}
