package main

import (
	"aluraBank/accounts"
	"aluraBank/clients"
	"fmt"
)

func main() {

	accountBruno := accounts.CheckingAccount{Owner: clients.Owner{Name: "Bruno", CPF: "123.456.789-00", Job: "Dev"}, AgencyNumber: 3123, AccountNumber: 312}
	accountBruno.Deposit(100000000)
	fmt.Println(accountBruno)
	fmt.Println(accountBruno.Getbalance())
	// accountMaria := accounts.CheckingAccount{Owner: "Maria", Balance: 100}
	// accountJoao := accounts.CheckingAccount{Owner: "Joao", Balance: 200}

	// status := accountJoao.Tranfer(-200, &accountMaria)

	// fmt.Println(status)
	// fmt.Println(accountMaria)
	// fmt.Println(accountJoao)
}
