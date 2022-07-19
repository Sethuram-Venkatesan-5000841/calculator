package calculator

type Calculator struct {
	result float64
}

func NewCalculator() Calculator {
	return Calculator{}
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
}
