package accounts

import "github.com/rkweber-max/go-object-oriented/bank/customers"

type SavingsAccount struct {
	Owner                                  customers.Customer
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (c *SavingsAccount) Withdraw(amountWithdraw float64) string {
	canWithdraw := amountWithdraw > 0 && amountWithdraw <= c.balance
	if canWithdraw {
		c.balance -= amountWithdraw
		return "Withdrawal successful"
	} else {
		return "Insufficient balance"
	}
}

func (c *SavingsAccount) Deposit(amountDeposit float64) (string, float64) {
	canDeposit := amountDeposit > 0
	if canDeposit {
		c.balance += amountDeposit
		return "Deposit successful", c.balance
	} else {
		return "Invalid deposit amount", c.balance
	}
}

func (s *SavingsAccount) GetBalance() float64 {
	return s.balance
}
