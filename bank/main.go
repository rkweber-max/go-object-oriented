package main

import (
	"fmt"

	"github.com/rkweber-max/go-object-oriented/bank/accounts"
)

func paymentBoleto(account *accounts.CurrentAccount, amount float64) {
	account.Withdraw(amount)
}

type verifyAccount interface {
	Withdraw(amount float64) string
}

func main() {
	accountRodrigo := accounts.CurrentAccount{}
	accountRodrigo.Deposit(100.00)

	paymentBoleto(&accountRodrigo, 100.00)

	fmt.Println(accountRodrigo.GetBalance())
}
