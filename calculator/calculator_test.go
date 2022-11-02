package calculator_test

import (
	"fmt"
	"github.com/aguazul-marco/pivot/calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name   string
		inputA int
		inputB int
		want   int
	}{
		{"testOne", 5, 4, 9},
		{"testTwo", 5, 5, 10},
		{"testThree", 7, 8, 15},
		{"testFour", 2, 4, 6},
		{"zeros", 0, 0, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := calculator.Add(test.inputA, test.inputB)
			if got != test.want {
				t.Errorf("got %q, wanted %q", got, test.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name   string
		inputA int
		inputB int
		want   int
	}{
		{"testOne", 5, 4, 1},
		{"testTwo", 5, 5, 0},
		{"testThree", 546, 230, 316},
		{"testFour", 2, 4, -2},
		{"zeros", 0, 0, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := calculator.Subtract(test.inputA, test.inputB)
			if got != test.want {
				t.Errorf("got %q, wanted %q", got, test.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name   string
		inputA int
		inputB int
		want   int
	}{
		{"testOne", 5, 4, 20},
		{"testTwo", 5, 5, 25},
		{"testThree", 546, 230, 125580},
		{"testFour", 2, 4, 8},
		{"zeros", 0, 0, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := calculator.Multiply(test.inputA, test.inputB)
			if got != test.want {
				t.Errorf("got %q, wanted %q", got, test.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name   string
		inputA int
		inputB int
		want   int
	}{
		{"testOne", 20, 4, 5},
		{"testTwo", 25, 5, 5},
		{"testThree", 546, 230, 2},
		{"testFour", 64, 4, 16},
		{"zeros", 0, 0, 0},
		{"divByZeroB", 8, 0, 0},
		{"divByZeroA", 0, 8, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got, err := calculator.Divide(test.inputA, test.inputB); err != nil {
				fmt.Println(err)
			} else if got != test.want {
				t.Errorf("got %q, wanted %q", got, test.want)
			}
		})
	}
}
