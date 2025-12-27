package main

import "fmt"

// Declaramos o nome dos valores que serão retornardos
func calcMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return

}
func main() {
	valor1, valor2 := calcMatematicos(3, 2)
	fmt.Println(valor1, valor2)
}
