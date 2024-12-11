package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {

	t.Run("Blink", func(t *testing.T) {
		assert.Equal(t, 3, sum("125 17", 1, &Cache{}))
		assert.Equal(t, 4, sum("125 17", 2, &Cache{}))
		assert.Equal(t, 5, sum("125 17", 3, &Cache{}))
		assert.Equal(t, 9, sum("125 17", 4, &Cache{}))
		assert.Equal(t, 13, sum("125 17", 5, &Cache{}))
		assert.Equal(t, 22, sum("125 17", 6, &Cache{}))
		assert.Equal(t, 55312, sum("125 17", 25, &Cache{}))
	})

}
