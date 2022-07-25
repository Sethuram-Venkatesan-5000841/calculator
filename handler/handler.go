package handler

import (
	"calculator/calculator"
	"calculator/models"
	"calculator/renderer"
)

type function func(operand float64, calcInterface calculator.ICalculator)

type FunctionMap map[models.Operations]function

var Registry = make(FunctionMap)

var commands map[models.Operations]float64

func RegisterHandler(operationString models.Operations, operation function) bool {
	for op, _ := range Registry {
		if operationString == op {
			panic("Duplicate entry")
		}
	}
	Registry[operationString] = operation
	return false
}

func ExecuteHandler(operand float64, operationString models.Operations, calcInterface calculator.ICalculator) {
	if Registry[operationString] == nil {
		renderer.ViewError("Invalid Function")
		return
	}

	Registry[operationString](operand, calcInterface)
}

func HandleAddition(operand float64, calcInterface calculator.ICalculator) {
	calcInterface.Add(operand)
	renderer.ViewValue(calcInterface.GetResults())
}

func HandleSubtraction(operand float64, calcInterface calculator.ICalculator) {
	calcInterface.Subtract(operand)
	renderer.ViewValue(calcInterface.GetResults())
}

func HandleMultiplication(operand float64, calcInterface calculator.ICalculator) {
	calcInterface.Multiply(operand)
	renderer.ViewValue(calcInterface.GetResults())
}

func HandleDivision(operand float64, calcInterface calculator.ICalculator) {
	calcInterface.Divide(operand)
	renderer.ViewValue(calcInterface.GetResults())
}

func HandleCancel(operand float64, calcInterface calculator.ICalculator) {
	calcInterface.Cancel()
	renderer.ViewValue(calcInterface.GetResults())
}

func HandleExit(operand float64, calcInterface calculator.ICalculator) {
	calcInterface.Exit()
}

func init() {
	RegisterHandler("add", HandleAddition)
	RegisterHandler("subtract", HandleSubtraction)
	RegisterHandler("multiply", HandleMultiplication)
	RegisterHandler("divide", HandleDivision)
	RegisterHandler("cancel", HandleCancel)
	RegisterHandler("exit", HandleExit)
}
