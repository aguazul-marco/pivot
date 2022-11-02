package calculator

import "errors"

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
