package renderer

import (
	"fmt"
)

func ViewValue(value float64) {
	fmt.Println(value)
}

func ViewError(error string) {
	fmt.Println(error)
}

func ViewSymbol() {
	fmt.Print("> ")
}
