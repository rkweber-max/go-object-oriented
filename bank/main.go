package main

import "fmt"

type CurrentAccount struct {
	owner         string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func (c *CurrentAccount) Withdraw(amountWithdraw float64) string {
	canWithdraw := amountWithdraw > 0 && amountWithdraw <= c.balance
	if canWithdraw {
		c.balance -= amountWithdraw
		return "Withdrawal successful"
	} else {
		return "Insufficient balance"
	}
}

func main() {
	accountRodrigo := CurrentAccount{
		owner:         "Rodrigo",
		agencyNumber:  123,
		accountNumber: 123456,
		balance:       1234.50,
	}

	fmt.Println("Balance Rodrigo:", accountRodrigo.balance)
	fmt.Println(accountRodrigo.Withdraw(100.00))
	fmt.Println("Balance Rodrigo after withdrawal:", accountRodrigo.balance)
}
