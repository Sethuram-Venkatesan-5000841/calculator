package reader

import (
	"bufio"
	"io"
	"os"
)

func Read(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	text := scanner.Text()
	if text == "exit" {
		defer func() { os.Exit(0) }()
	}
	if text == "cancel" {
		text = "multiply 0"
	}
	return text
}
