package main

import (
	"aluraBank/accounts"
	"fmt"
)

func main() {
	accountMaria := accounts.CheckingAccount{Owner: "Maria", Balance: 100}
	accountJoao := accounts.CheckingAccount{Owner: "Joao", Balance: 200}

	status := accountJoao.Tranfer(-200, &accountMaria)

	fmt.Println(status)
	fmt.Println(accountMaria)
	fmt.Println(accountJoao)
}
