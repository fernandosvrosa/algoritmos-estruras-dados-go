package cap4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividirConquistarRecursao(t *testing.T) {
	t.Run("Dividir para conquistar 12", func(t *testing.T) {
		resultado := soma([]int{2, 4, 6})
		esperado := 12

		assert.Equal(t, esperado, resultado, "Test soma de 2, 4 e 6 = 12")

	})

	t.Run("Dividir para conquistar 14", func(t *testing.T) {
		resultado := soma([]int{2, 4, 6, 2})
		esperado := 14

		assert.Equal(t, esperado, resultado, "Test soma de 2, 4 , 6, 2 = 14")

	})

	t.Run("Dividir para conquistar 0", func(t *testing.T) {
		resultado := soma([]int{})
		esperado := 0

		assert.Equal(t, esperado, resultado, "Test soma de 0 = 0")

	})

}
