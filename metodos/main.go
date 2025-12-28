package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuario: %s", u.nome)
}

func (u *usuario) fazerAniversario() {
	u.idade++
}
func main() {
	usuario1 := usuario{"mello", 20}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Teste", 2}
	usuario2.salvar()
}
