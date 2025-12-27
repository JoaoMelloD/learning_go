package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	idade := 10
	numero := 20

	if idade >= 10 {
		fmt.Println("Idade maior que 10")
	} else {
		fmt.Println("Idade menor que 10")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que zero")
	}
}
