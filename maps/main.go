package main

import "fmt"

func main() {

	fmt.Println("Maps")

	// Dentro do colchetes => Tipos das chaves, fora deles é o tipo do valor
	usuario := map[string]string{
		"nome":      "Mello",
		"sobrenome": "Delgado",
	}

	fmt.Println(usuario)

	// Maps Aninhados, ou seja, inserindo um map dentro de outro map no qual a chave
	// dentro do map dentro do outro map no qual a key é string e o value também
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Joao",
			"segundo":  "mello",
		},
		"endereco": {
			"rua":    "montes claros",
			"bairro": "jardim global",
		},
	}

	// Funcao usada para deletar uma chave dentro do map
	delete(usuario2, "endereco")

	fmt.Println(usuario2)
}
