package nreinas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func esSolucion(solucion []int) bool {
	for i := range solucion {
		for j := i + 1; j < len(solucion); j++ {
			if solucion[i] == solucion[j] ||
				j-i == solucion[j]-solucion[i] ||
				j-i == solucion[i]-solucion[j] {
				return false
			}
		}
	}
	return true
}

func TestNReinas(t *testing.T) {
	n := 12

	soluciones := NReinas(n)

	for _, sol := range soluciones {
		assert.Len(t, sol, n)
		assert.True(t, esSolucion(sol))
	}
}