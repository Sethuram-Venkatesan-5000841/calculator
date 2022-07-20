package parser

import (
	"strconv"
	"strings"
)

type Parser struct {
	tokens []string
}

func NewParser() Parser {
	return Parser{}
}

func (parser *Parser) Parse(input string) (string, float64) {
	if input == "" {
		panic("cannot parse empty string")
	}

	parser.tokens = strings.Split(input, " ")

	if len(parser.tokens) != 2 {
		panic("invalid input")
	}

	operand, err := strconv.ParseFloat(parser.tokens[1], 64)
	if err != nil {
		panic("invalid operand")
	}

	return parser.tokens[0], operand
}
