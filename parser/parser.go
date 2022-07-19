package parser

type Parser struct {
	tokens []string
}

func NewParser() Parser {
	return Parser{}
}

func (parser Parser) Parse(input string) {
	if input == "" {
		panic("cannot parse empty string")
	}
}
