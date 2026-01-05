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

func (c *CurrentAccount) Deposit(amountDeposit float64) (string, float64) {
	canDeposit := amountDeposit > 0
	if canDeposit {
		c.balance += amountDeposit
		return "Deposit successful", c.balance
	} else {
		return "Invalid deposit amount", c.balance
	}
}

func (c *CurrentAccount) Transfer(amountTransfer float64, destinationAccount *CurrentAccount) string {
	canTransfer := amountTransfer > 0 && amountTransfer <= c.balance
	if canTransfer {
		c.balance -= amountTransfer
		destinationAccount.balance += amountTransfer
		return "Transfer successful"
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

	accountAmanda := CurrentAccount{
		owner:         "Amanda",
		agencyNumber:  123,
		accountNumber: 123456,
		balance:       1234.50,
	}

	fmt.Println("Balance Rodrigo:", accountRodrigo.balance)
	fmt.Println(accountRodrigo.Transfer(100.00, &accountAmanda))
	fmt.Println("Balance Rodrigo after transfer:", accountRodrigo.balance)
	fmt.Println("Balance Amanda after transfer:", accountAmanda.balance)
}
