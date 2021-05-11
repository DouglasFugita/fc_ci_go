package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da Soma invalido: Resultado %d, Valor esperado %d", total, 30)
	}
}
