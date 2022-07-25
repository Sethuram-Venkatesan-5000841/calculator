package main

import (
	"calculator/calculator"
	"calculator/handler"
	"calculator/models"
	"calculator/parser"
	"calculator/reader"
	"calculator/renderer"
	"os"
)

func main() {
	calci := calculator.NewCalculator()
	for true {
		renderer.ViewSymbol()
		Input := reader.Read(os.Stdin)
		parse := parser.NewParser()
		operator, operand := parse.Parse(Input)
		handler.ExecuteHandler(operand, models.Operations(operator), calci)
	}
}
