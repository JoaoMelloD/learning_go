package main

import "fmt"

func somar(a int8, b int8) int8 {
	return a + b
}

var f = func(txt string) string {
	return txt
}

// Recebe dois parametros int8 mas da dois retornos
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	y := somar(30, 100)
	fmt.Println(y)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(1, 2)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}
