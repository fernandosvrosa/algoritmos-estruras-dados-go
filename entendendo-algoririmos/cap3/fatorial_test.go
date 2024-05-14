package cap3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFatorial(t *testing.T) {
	t.Run("Fatorial de 5", func(t *testing.T) {
		numero := 5
		resultado := fatorial(numero)
		esperado := 120

		assert.Equal(t, resultado, esperado)

	})

	t.Run("Fatorial de 0", func(t *testing.T) {
		numero := 0
		resultado := fatorial(numero)
		esperado := 1

		assert.Equal(t, resultado, esperado, "Test fatorial 0")
	})
}
