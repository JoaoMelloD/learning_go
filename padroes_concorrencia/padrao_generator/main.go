package main

import (
	"fmt"
	"time"
)

func main() {

	// Quando for chamar minha funcao nao preciso realizar o uso do prefixo Go
	// Minha funcao encapsula a chamada a go routine e retorna um chanel
	canal := escrever("Olá Mundo")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}
func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("Valor Recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
