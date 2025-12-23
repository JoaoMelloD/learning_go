package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Mello")
	esperado := "Olá, Mello"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", esperado, resultado)
	}
}
