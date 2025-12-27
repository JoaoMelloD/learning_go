package main

import "fmt"

func funcao1() {
	fmt.Println("Escrevendo Funcao 1")
}

func funcao2() {
	fmt.Println("Escrevendo Funcao 2")
}

func main() {
	defer funcao2()

	funcao1()
}
