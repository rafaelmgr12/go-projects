package main

import "fmt"

func main() {
	// var nome string = "Rafael"
	// var idade int = 27
	// var versao float32 = 1.2

	// go can infer variables types

	// var nome = "Rafael"
	// var idade = 27
	// var versao = 1.2

	// go can also take out the var name and still work

	nome := "Rafael"
	idade := 27
	versao := 1.2

	fmt.Println("Olá, sr.", nome, "sua	idade é", idade)
	fmt.Println("Este programa está na versão", versao)
}
