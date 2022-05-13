package main

import "fmt"

func main() {
	name := "Rafael"
	version := 1.1

	fmt.Println("Hello, ", name, "!")
	fmt.Println("The program is in the version ", version)

	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show Logs")
	fmt.Println("3- Exit")

	var option int
	//fmt.Scanf("%d", &option)
	// it is possible to use anothe way to read the option,
	//and it is not necessary to explictly convert the type
	fmt.Scan(&option)

	fmt.Println("You selected option ", option)
}
