package calculator

import (
	"errors"
	"math"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("can't divide a larger denominator by zero")
	}
	return a / b, nil
}

func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}
