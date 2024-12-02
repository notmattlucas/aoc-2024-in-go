package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollections(t *testing.T) {

	t.Run("Any", func(t *testing.T) {
		actual, idx := Any([]int{1, 2, 3, 4, 5}, func(x int) bool {
			return x == 3
		})
		assert.True(t, actual)
		assert.Equal(t, 2, idx)

		actual, idx = Any([]int{1, 2, 3, 4, 5}, func(x int) bool {
			return x == 7
		})
		assert.False(t, actual)
	})

	t.Run("All", func(t *testing.T) {
		actual, idx := All([]int{1, 2, 3, 4, 5}, func(x int) bool {
			return x > 0
		})
		assert.True(t, actual)
		assert.Equal(t, -1, idx)

		actual, idx = All([]int{1, 2, 3, 4, 5}, func(x int) bool {
			return x > 1
		})
		assert.False(t, actual)
		assert.Equal(t, 0, idx)
	})

}
