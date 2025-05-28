package cambio

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCambio(t *testing.T) {
	monedas := []int{1, 2, 9, 10, 20, 50}
	cantidad := 27

	solucion := CambioDeMoneda(cantidad, monedas)

	sum := 0
	for _, coin := range solucion {
		sum += coin
	}
	require.Equal(t, cantidad, sum)
	require.Len(t, solucion, 3) // 9 * 3 = 27
	for _, coin := range solucion {
		require.Equal(t, 9, coin)
	}
}
