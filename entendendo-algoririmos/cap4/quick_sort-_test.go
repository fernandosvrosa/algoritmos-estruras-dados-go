package cap4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSortPivoIndexZero(t *testing.T) {
	t.Run("array  size 1", func(t *testing.T) {
		resultado := quickSortPivoIndexZero([]int{2})
		esperado := []int{2}

		assert.Equal(t, esperado, resultado, "size 1")

	})

	t.Run("Array ordenado", func(t *testing.T) {
		resultado := quickSortPivoIndexZero([]int{5, 2, 4, 6, 3})
		esperado := []int{2, 3, 4, 5, 6}

		assert.Equal(t, esperado, resultado, "Array ordenado")

	})

}
