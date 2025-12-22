package main

import "fmt"

// Ola retorna uma saudação padrão.
func Ola(nome string) string {
	return "Olá, " + nome
}

func GoodBye(nome string) string {
	return "Até Logo " + nome
}

func main() {
	fmt.Println("Olá Mundo")
}
