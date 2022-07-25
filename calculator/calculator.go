package calculator

import "os"

type ICalculator interface {
	Add(number float64)
	Subtract(number float64)
	Multiply(number float64)
	Divide(number float64)
	GetResults() float64
	Cancel()
	Exit()
}

type Calculator struct {
	result float64
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (calculator *Calculator) Add(number float64) {
	calculator.result += number
}

func (calculator *Calculator) Subtract(number float64) {
	calculator.result -= number
}

func (calculator *Calculator) Multiply(number float64) {
	calculator.result *= number
}

func (calculator *Calculator) Divide(number float64) {
	if number == 0 {
		panic("Cannot divide by 0!")
	}
	calculator.result /= number
}
func (calculator *Calculator) GetResults() float64 {
	return calculator.result
}

func (calculator *Calculator) Cancel() {
	calculator.result = 0
}

func (calculator *Calculator) Exit() {
	os.Exit(0)
}
