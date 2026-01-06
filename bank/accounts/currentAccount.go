package accounts

import "github.com/rkweber-max/go-object-oriented/bank/customers"

type CurrentAccount struct {
	Owner         customers.Customer
	AgencyNumber  int
	AccountNumber int
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
	canTransfer := amountTransfer > 0 && amountTransfer <= c.balance && c.balance > 0 && destinationAccount != nil
	if canTransfer {
		c.balance -= amountTransfer
		destinationAccount.balance += amountTransfer
		return "Transfer successful"
	} else {
		return "Insufficient balance"
	}
}

func (c *CurrentAccount) GetBalance() float64 {
	return c.balance
}
