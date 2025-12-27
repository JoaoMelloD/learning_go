package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)
	//teve seu valor aumentado
	variavel1++

	fmt.Println(variavel1, variavel2)

	// Usando ponteiro, estamos usando uma referência de memoria
	var variavelA int
	//guarda o end de memoria de um inteiro
	var ponteiro *int

	fmt.Println(variavelA, ponteiro)

	variavelA = 100
	ponteiro = &variavelA // usando o & ele avisa que esta apontando para o end de memoria da variavelA

	//Quando usamos * na declaração do ponteiro ele nada mais que pega o end da memoria do valor que o ponteiro mostra
	//e apos isso resgata o valor armazenado, chamado de desreferênciação
	fmt.Println(variavelA, *ponteiro)

	variavelA = 301
	fmt.Println(variavelA, *ponteiro)

	var valorRef int = 103
	var ponteiro3 *int = &valorRef
	var ponteiro2 *int = ponteiro3

	println(valorRef, *ponteiro2, *ponteiro3)

	valorRef = 9
	println(valorRef, *ponteiro2, *ponteiro3)

}
