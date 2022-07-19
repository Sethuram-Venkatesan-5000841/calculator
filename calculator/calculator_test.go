package calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCalculator(t *testing.T) {
	t.Run("when initializing calculator, calculator entity should be returned", func(t *testing.T) {
		assert.IsType(t, Calculator{}, NewCalculator())
	})

	t.Run("when initializing calculator the initial value should be 0.0", func(t *testing.T) {
		assert.Equal(t, 0.0, NewCalculator().result)
	})
}
