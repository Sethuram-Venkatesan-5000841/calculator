package reader

import (
	"bufio"
	"io"
)

func Read(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	return scanner.Text()
}
