package main

import "testing"

func TestSum(t *testing.T) {
	total := sum(15, 15)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}

func TestSubtract(t *testing.T) {	
	total := subtract(15, 5)

	if total != 10 {
		t.Errorf("Resultado da subtração é inválido: Resultado %d. Esperado: %d", total, 10)
	}
}

func TestMultiply(t *testing.T) {
	total := multiply(5, 5)

	if total != 25 {
		t.Errorf("Resultado da multiplicação é inválido: Resultado %d. Esperado: %d", total, 25)
	}
}

func TestDivide(t *testing.T) {
	total := divide(10, 2)

	if total != 5 {
		t.Errorf("Resultado da divisão é inválido: Resultado %d. Esperado: %d", total, 5)
	}
}

func TestDivideByZero(t *testing.T) {
	total := divide(10, 0)

	if total != 0 {
		t.Errorf("Resultado da divisão por zero é inválido: Resultado %d. Esperado: %d", total, 0)
	}
}

func TestModulo(t *testing.T) {
	total := modulo(10, 3)

	if total != 1 {
		t.Errorf("Resultado do módulo é inválido: Resultado %d. Esperado: %d", total, 1)
	}
}

func TestModuloByZero(t *testing.T) {
	total := modulo(10, 0)

	if total != 0 {
		t.Errorf("Resultado do módulo por zero é inválido: Resultado %d. Esperado: %d", total, 0)
	}
}