package main

import "fmt"

// Estou indicando que a funcao vai receber de 1 a N int
func soma(numeros ...int) {
	total := 0
	for _, valor := range numeros {
		total += valor
	}
	println(total)
}

// Só é possivel ter um valor variatico por função e ele sempre deve ser declarado por ultimo
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	escrever("Teste: ", 1, 2, 3, 4, 5, 6, 7)

}
