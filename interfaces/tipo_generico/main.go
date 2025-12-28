package main

import "fmt"

// Nos Permite interagir com qualquer coisa
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(false)
}
