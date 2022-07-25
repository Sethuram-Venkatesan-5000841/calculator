package parser

import (
	"strconv"
	"strings"
)

type Parser struct {
}

func NewParser() Parser {
	return Parser{}
}

func (parser *Parser) Parse(input string) (string, float64) {
	if input == "" {
		panic("cannot parse empty string")
	}

	tokens := strings.Split(input, " ")

	if len(tokens) != 2 {
		panic("invalid input")
	}

	operand, err := strconv.ParseFloat(tokens[1], 64)
	if err != nil {
		panic("invalid operand")
	}

	return tokens[0], operand
}
