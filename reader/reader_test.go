package reader

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	var reader *strings.Reader
	t.Run("while getting input from the user the reader should return string", func(t *testing.T) {
		reader = strings.NewReader("add 5")
		assert.IsType(t, "", Read(reader))
	})

	t.Run("while getting input from user, returned string should match the user input", func(t *testing.T) {
		reader = strings.NewReader("sub 10")
		assert.Equal(t, "sub 10", Read(reader))
	})

	t.Run("while getting input from user, returned string should match the user input", func(t *testing.T) {
		reader = strings.NewReader("multiply 10")
		assert.Equal(t, "multiply 10", Read(reader))
	})

	t.Run("should exit the program when 'exit' command is passed by the user", func(t *testing.T) {
		reader = strings.NewReader("exit")
		assert.Panics(t, func() {
			Read(reader)
		})
	})
}
