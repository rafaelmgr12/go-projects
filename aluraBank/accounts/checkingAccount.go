package accounts

type CheckingAccount struct {
	Owner         string
	AccountNumber int
	AgencyNumber  int
	Balance       float64
}

func (c *CheckingAccount) Withdraw(valueWithdraw float64) string {
	canWithdraw := valueWithdraw > 0 && valueWithdraw <= c.Balance

	if canWithdraw {
		c.Balance -= valueWithdraw
		return "Success Withdraw"
	} else {
		return "Insufficient funds"
	}
}

func (c *CheckingAccount) Deposit(valueDeposit float64) string {
	if valueDeposit > 0 {
		c.Balance += valueDeposit
		return "Success Deposit"
	} else {
		return "Invalid value"
	}
}

func (c *CheckingAccount) Tranfer(valueTranfer float64, accountDestiny *CheckingAccount) string {
	if valueTranfer > 0 && valueTranfer <= c.Balance {
		c.Balance -= valueTranfer
		accountDestiny.Deposit(valueTranfer)
		return "Success Transfer"
	} else {
		return "Insufficient funds"
	}
}
