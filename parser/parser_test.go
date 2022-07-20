package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewParser(t *testing.T) {
	t.Run("when initializing parser, should return a parser", func(t *testing.T) {
		assert.IsType(t, Parser{}, NewParser())
	})
}

func TestParser_Parse(t *testing.T) {
	var parser Parser = NewParser()
	t.Run("when parsing the string should not be empty", func(t *testing.T) {
		assert.Panics(t, func() {
			parser.Parse("")
		})
	})

	t.Run("should panic when the input string does not have two tokens", func(t *testing.T) {
		assert.Panics(t, func() {
			parser.Parse("add")
		})

		assert.Panics(t, func() {
			parser.Parse("add sub 5")
		})
	})

	t.Run("should return a string and a float value for the input 'divide 5.5'", func(t *testing.T) {
		operator, operand := parser.Parse("divide 5.5")
		assert.IsType(t, "", operator)
		assert.IsType(t, 0.0, operand)
	})

	t.Run("should return ['add', 5.0] for the input 'add 5'", func(t *testing.T) {
		operator, operand := parser.Parse("add 5")
		assert.Equal(t, "add", operator)
		assert.Equal(t, 5.0, operand)
	})

	t.Run("should return ['subtract', 10.0] for the input 'subtract 10'", func(t *testing.T) {
		operator, operand := parser.Parse("subtract 10")
		assert.Equal(t, "subtract", operator)
		assert.Equal(t, 10.0, operand)
	})

	t.Run("should return ['multiply', 3.0] for the input 'multiply 3'", func(t *testing.T) {
		operator, operand := parser.Parse("multiply 3")
		assert.Equal(t, "multiply", operator)
		assert.Equal(t, 3.0, operand)
	})

	t.Run("should return ['divide', 2.5] for the input 'divide 2.5'", func(t *testing.T) {
		operator, operand := parser.Parse("divide 2.5")
		assert.Equal(t, "divide", operator)
		assert.Equal(t, 2.5, operand)
	})
}
