package main

import (
	"fmt"
	"os"
)

func main() {

	showInfo()
	showMenu()
	option := readOption()

	// if option == 1 {
	// 	fmt.Println("Start Monitoring")
	// } else if option == 2 {
	// 	fmt.Println("Show Logs")
	// } else if option == 0 {
	// 	fmt.Println("Exit")
	// } else {
	// 	fmt.Println("Invalid Option")
	// }
	// it is possible to use switch case in Go
	switch option {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Show Logs...")
	case 0:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid Option")
		os.Exit(-1)
	}

}

func showInfo() {
	name := "Rafael"
	version := 1.1

	fmt.Println("Hello, ", name, "!")
	fmt.Println("The program is in the version ", version)
}

func showMenu() {
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")

}

func readOption() int {
	var option int
	//fmt.Scanf("%d", &option)
	// it is possible to use anothe way to read the option,
	//and it is not necessary to explictly convert the type
	fmt.Scan(&option)

	fmt.Println("You selected option ", option)
	return option
}
