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

func TestAdd(t *testing.T) {
	t.Run("should return 0.0 for number 0.0 when the result is 0.0", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(0.0)
		assert.Equal(t, 0.0, calculator.result)
	})

	t.Run("should return 5.0 for number 5.0 when the result is 0.0", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(5)
		assert.Equal(t, 5.0, calculator.result)
	})

	t.Run("should return -5.0 for number -5.0 when the result is 0.0", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(-5.0)
		assert.Equal(t, -5.0, calculator.result)
	})

	t.Run("should return 5.5 for number 5.5 when the result is 0.0", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(5.5)
		assert.Equal(t, 5.5, calculator.result)
	})

	t.Run("should return 11.0 for number 5.5 when the result is 5.5", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(5.5)
		calculator.Add(5.5)
		assert.Equal(t, 11.0, calculator.result)
	})

	t.Run("should return 0.0 for number 5.5 when the result is -5.5", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(-5.5)
		calculator.Add(5.5)
		assert.Equal(t, 0.0, calculator.result)
	})

	t.Run("should return -11.0 for number -5.5 when the result is -5.5", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(-5.5)
		calculator.Add(-5.5)
		assert.Equal(t, -11.0, calculator.result)
	})
}

func TestSubtract(t *testing.T) {
	t.Run("should return 0.0 for number 0.0 when the result is 0.0", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Subtract(0)
		assert.Equal(t, 0.0, calculator.result)
	})

	t.Run("should return -5.0 for number 5.0 when the result is 0.0", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Subtract(5)
		assert.Equal(t, -5.0, calculator.result)
	})

	t.Run("should return 5.0 for number -5.0 when the result is 0.0", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Subtract(-5)
		assert.Equal(t, 5.0, calculator.result)
	})

	t.Run("should return 6.5 for number 3.5 when the result is 10.0", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(10)
		calculator.Subtract(3.5)
		assert.Equal(t, 6.5, calculator.result)
	})
}
