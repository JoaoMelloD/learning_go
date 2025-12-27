package main

import "fmt"

func main() {
	// fmt.Println("Loops dentro do Go")

	nomes := []string{"Joao", "Mello", "Delgado"}

	// for i := 0; i < len(nomes); {
	// 	fmt.Println("Nome: ", nomes[i])
	// 	i++
	// }

	for indice, nome := range nomes {
		fmt.Print("Nome: ", nome)
		fmt.Println(" - Posicao: ", indice)

	}
}
