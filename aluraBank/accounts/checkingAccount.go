package accounts

import (
	"aluraBank/clients"
)

type CheckingAccount struct {
	Owner         clients.Owner
	AccountNumber int
	AgencyNumber  int
	balance       float64
}

func (c *CheckingAccount) Withdraw(valueWithdraw float64) string {
	canWithdraw := valueWithdraw > 0 && valueWithdraw <= c.balance

	if canWithdraw {
		c.balance -= valueWithdraw
		return "Success Withdraw"
	} else {
		return "Insufficient funds"
	}
}

func (c *CheckingAccount) Deposit(valueDeposit float64) string {
	if valueDeposit > 0 {
		c.balance += valueDeposit
		return "Success Deposit"
	} else {
		return "Invalid value"
	}
}

func (c *CheckingAccount) Tranfer(valueTranfer float64, accountDestiny *CheckingAccount) string {
	if valueTranfer > 0 && valueTranfer <= c.balance {
		c.balance -= valueTranfer
		accountDestiny.Deposit(valueTranfer)
		return "Success Transfer"
	} else {
		return "Insufficient funds"
	}
}

func (c *CheckingAccount) Getbalance() float64 {
	return c.balance
}
