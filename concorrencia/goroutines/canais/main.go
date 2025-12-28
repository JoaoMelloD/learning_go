package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá Mundo", canal)

	// Como está em loop infinito, o channel fica esperando infinito, entrando em deadlock.
	// Deadlock só da erro em execução
	// É possível verificar se o canal está aberto ou fechado para evitar deadlocks.
	for {
		// Esperando que o canal receba um valor
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)

	}

	fmt.Println("Fim Do Programa")
}

// Deadlock -> Quando nao tem mais nenhum lugar enviando dados ao canal mas o canal ainda está esperando dados.

func escrever(texto string, canal chan string) {

	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
