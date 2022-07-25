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

func TestMultiply(t *testing.T) {
	t.Run("should return 0.0 for number 0.0 when the result is 0.0", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Multiply(0)
		assert.Equal(t, 0.0, calculator.result)
	})

	t.Run("should return 10.0 for number 5.0 when the result is 2.0", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(2)
		calculator.Multiply(5)
		assert.Equal(t, 10.0, calculator.result)
	})

	t.Run("should return -10.0 for number -5.0 when the result is 2.0", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(2)
		calculator.Multiply(-5)
		assert.Equal(t, -10.0, calculator.result)
	})

	t.Run("should return 10.0 for number -5.0 when the result is -2.0", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(-2)
		calculator.Multiply(-5)
		assert.Equal(t, 10.0, calculator.result)
	})

	t.Run("should return 7.0 for number 2.0 when the result is 3.5", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(3.5)
		calculator.Multiply(2.0)
		assert.Equal(t, 7.0, calculator.result)
	})

	t.Run("should return -7.0 for number -2 when the result is 3.5", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(3.5)
		calculator.Multiply(-2)
		assert.Equal(t, -7.0, calculator.result)
	})
}

func TestDivide(t *testing.T) {
	t.Run("should raise error when dividied by zero", func(t *testing.T) {
		calculator := NewCalculator()
		assert.Panics(t, func() {
			calculator.Divide(0.0)
		})
	})

	t.Run("should return 5.0 for number 2.0 when the result is 10.0", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(10)
		calculator.Divide(2.0)
		assert.Equal(t, 5.0, calculator.result)
	})

	t.Run("should return -5.0 for number 2.0 when the result is -10.0", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Subtract(10)
		calculator.Divide(2.0)
		assert.Equal(t, -5.0, calculator.result)
	})

	t.Run("should return 5.0 for number -2.0 when the result is -10.0", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Subtract(10)
		calculator.Divide(-2.0)
		assert.Equal(t, 5.0, calculator.result)
	})

	t.Run("should return 2.5 for number 2.0 when the result is 5.0", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(5)
		calculator.Divide(2.0)
		assert.Equal(t, 2.5, calculator.result)
	})

	t.Run("should return 2.0 for number 2.5 when the result is 5.0", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(5)
		calculator.Divide(2.5)
		assert.Equal(t, 2.0, calculator.result)
	})

	t.Run("should return 0.5 for number 5.0 when the result is 2.5", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(2.5)
		calculator.Divide(5.0)
		assert.Equal(t, 0.5, calculator.result)
	})
}

func TestCancel(t *testing.T) {
	t.Run("should return 0 when user enters cancel button", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(50)
		calculator.Cancel()
		assert.Equal(t, 0.0, calculator.result)
	})
}

func TestGetResults(t *testing.T) {
	t.Run("should return 0 when the calculator has the result as 0", func(t *testing.T) {
		calculator := NewCalculator()
		assert.Equal(t, 0.0, calculator.GetResults())
	})

	t.Run("should return 500 when the calculator has the result as 500", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(500)
		assert.Equal(t, 500.0, calculator.GetResults())
	})
}
