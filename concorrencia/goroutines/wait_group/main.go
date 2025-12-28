package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// Maneira de Sincronizar as goroutiness para que a função main seja encerrada somente quando as rotinas serem finalizadas.
	var waitGroup sync.WaitGroup

	//Coloco a quantia de goroutines que fazem parte do grupo de espera
	waitGroup.Add(2)

	go func() {
		escrever("Olá Mundo")
		// Vai remover do contador 1 quando finalizar
		waitGroup.Done()
	}()

	go func() {
		escrever("Outro Texto")
		waitGroup.Done()
	}()

	// Fala pra função main esperar a contagem das goroutines serem finalizadas
	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
