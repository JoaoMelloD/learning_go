package main

import "fmt"

// Falando de Estruturas
type usuario struct {
	nome  string
	idade uint8
}

// Falando de Herança
type pessoa struct {
	nome   string
	idade  uint8
	altura uint8
}
type estudante struct {
	pessoa    // => Pegando todos os campos de pessoa e jogando aqui
	codigo    string
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Arquivo Sobre Structs")

	var u usuario
	u.nome = "Mello"
	u.idade = 20

	fmt.Println(u)

	usuario2 := usuario{
		nome:  "Mello",
		idade: 21,
	}
	fmt.Println(usuario2)

	pessoa1 := pessoa{
		nome:   "Mello",
		idade:  20,
		altura: 120,
	}

	fmt.Println(pessoa1)

	aluno := estudante{
		pessoa1,
		"SSSDKASDOASDASJO",
		"Sistemas De Informação",
		"Unipar",
	}

	fmt.Println(aluno)
}
