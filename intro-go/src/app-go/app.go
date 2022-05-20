package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	showInfo()

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
	for {
		showMenu()

		option := readOption()

		switch option {
		case 1:
			starMonitoring()
		case 2:
			fmt.Println("Show Logs...")
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
			break
		default:
			fmt.Println("Invalid Option")
			os.Exit(-1)
			break
		}
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

func starMonitoring() {
	monitoringTimes := 5
	delay := 10 * time.Minute // 10 minutes
	fmt.Println("Start Monitoring...")
	sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br/", "https://www.caelum.com.br/"}

	// for i := 0; i < len(sites); i++ {
	// 	fmt.Println("Sites: ", sites[i])

	// }
	for i := 0; i < monitoringTimes; i++ {
		for i, site := range sites {
			fmt.Println("Position in a slice is ", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay)
	}

}

func testSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site is up!")
	} else {
		fmt.Println("Site is down! Status code: ", resp.StatusCode)

	}

}
