package main

import "fmt"

func inverteSinal(numero int) int {
	return numero * -1
}

func inverteSinalPonteiro(numero *int) {
	*numero = *numero * -1

}

func main() {
	numero := 20
	// Não altera o valor da variável, cria uma nova var e faz uma copia desse valor e manda para função e dps a primeira var fica inalteravel
	numeroInvertido := inverteSinal(numero)
	fmt.Println(numeroInvertido)

	novoNumero := 40
	//Altera o valor da variavel usando o endereço de memoria por conta do ponteiro
	fmt.Println(novoNumero)
	inverteSinalPonteiro(&novoNumero)

	fmt.Println(novoNumero)

}
