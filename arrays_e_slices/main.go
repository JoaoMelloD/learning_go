package main

import "fmt"

func main() {

	fmt.Println("Arrays e Slices")
	//Arrays nada mais é do que uma lista de valores

	// Todos os dados do array devem ser do mesmo tipo.
	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]string{
		"Posicao1",
		"Posicao2",
		"Posicao3",
		"Posicao4",
		"Posicao5",
	}
	// Possuí um tamanho fixo e não é flexivel
	fmt.Println(array2)

	//Esses 3 pontinhos servem para fixar o tamanho do array com a quantidade de itens que ele recebe
	// Ou seja deixa o tamanho conforme a quantidade de valores que é atribuida nele
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//Slice
	slice := []int{10, 11, 2, 123, 3213, 1, 3, 13}
	fmt.Println(slice)
	//Possui uma função chamada append, que pega o slice, adiciona um valor novo, e retorna um slice novo com o mesmo nome.

	slice_dois := []int{10, 20, 30}
	fmt.Println(slice_dois)
	slice_dois = append(slice_dois, 40)
	fmt.Println(slice_dois)

	// Slice é uma fatia ou seja um pedaço de array
	// Pegou uma fatia a partir de um array que ja existia
	slice2 := array1[1:3]
	fmt.Println(slice2)
}
