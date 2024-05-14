package cap2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenacao(t *testing.T) {

	t.Run("Test list", func(t *testing.T) {
		result := ordenacaoPorSelecao([]int{5, 3, 6, 2, 10, 1})
		expected := []int{1, 2, 3, 5, 6, 10}

		assert.Equal(t, expected, result)
	})

}
