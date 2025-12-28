package main

func main() {
	canal := make(chan string, 2)
	canal <- "Pinto"
	canal <- "Penis"

	mensagem := <-canal
	mensagemB := <-canal

	println(mensagem, mensagemB)

}
