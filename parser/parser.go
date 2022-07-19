package parser

import "strings"

type Parser struct {
	tokens []string
}

func NewParser() Parser {
	return Parser{}
}

func (parser *Parser) Parse(input string) {
	if input == "" {
		panic("cannot parse empty string")
	}

	parser.tokens = strings.Split(input, " ")
}
