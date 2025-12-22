package main

import "testing"

func TestGoodbye(t *testing.T) {
	resultado := GoodBye("Jorge")
	esperado := "Até Logo Jorge"

	if resultado != esperado {
		t.Errorf("Resultado Errado '%s'", esperado)
	}
}
