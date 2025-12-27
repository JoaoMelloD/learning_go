package main

import (
	"fmt"
	"testes/functions"
)

type pessoa struct {
	nome  string
	idade int
}

func (p *pessoa) ascenar() string {
	return "Ola, meu nome e " + p.nome + " e tenho " + fmt.Sprint(p.idade) + " anos."
}

func main() {
	valorA := functions.Sum(1, 2)
	valorB := functions.Sub(2, 1)
	fmt.Println(valorB)
	var ponteiro *int = &valorA
	fmt.Println(valorA, ponteiro)

	
	usuario := pessoa{"Joao", 20}
	mensagem := usuario.ascenar()
	fmt.Println(mensagem)

}
