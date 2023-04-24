package mochila

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMochila6(t *testing.T) {
	objetos := []Objeto{
		{Peso: 1, Valor: 1},
		{Peso: 2, Valor: 6},
		{Peso: 3, Valor: 18},
		{Peso: 4, Valor: 22},
		{Peso: 5, Valor: 28},
	}
	capacidad := 11

	valor, _ := Mochila6(objetos, capacidad)

	assert.Equal(t, 56, valor)
}

func TestMochila6Bis(t *testing.T) {
	objetos := []Objeto{
		{Peso: 1, Valor: 10},
		{Peso: 2, Valor: 6},
		{Peso: 3, Valor: 18},
		{Peso: 4, Valor: 22},
		{Peso: 5, Valor: 28},
	}
	capacidad := 11

	valor, _ := Mochila6(objetos, capacidad)

	assert.Equal(t, 62, valor)
}
