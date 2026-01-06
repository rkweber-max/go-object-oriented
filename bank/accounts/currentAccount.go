package accounts

type CurrentAccount struct {
	Owner         string
	AgencyNumber  int
	AccountNumber int
	Balance       float64
}

func (c *CurrentAccount) Withdraw(amountWithdraw float64) string {
	canWithdraw := amountWithdraw > 0 && amountWithdraw <= c.Balance
	if canWithdraw {
		c.Balance -= amountWithdraw
		return "Withdrawal successful"
	} else {
		return "Insufficient balance"
	}
}

func (c *CurrentAccount) Deposit(amountDeposit float64) (string, float64) {
	canDeposit := amountDeposit > 0
	if canDeposit {
		c.Balance += amountDeposit
		return "Deposit successful", c.Balance
	} else {
		return "Invalid deposit amount", c.Balance
	}
}

func (c *CurrentAccount) Transfer(amountTransfer float64, destinationAccount *CurrentAccount) string {
	canTransfer := amountTransfer > 0 && amountTransfer <= c.Balance && c.Balance > 0 && destinationAccount != nil
	if canTransfer {
		c.Balance -= amountTransfer
		destinationAccount.Balance += amountTransfer
		return "Transfer successful"
	} else {
		return "Insufficient balance"
	}
}
