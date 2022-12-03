package main

import (
	"fmt"

	"github.com/aguazul-marco/pivot/calculator"
)

func main() {

	addResult := calculator.Add(8, 8)
	subResult := calculator.Subtract(8, 8)
	multiplyResult := calculator.Multiply(8, 8)
	if divideResult, err := calculator.Divide(8, 0); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(divideResult)
	}

	powResult := calculator.Pow(2, 3)

	fmt.Println(addResult)
	fmt.Println(subResult)
	fmt.Println(multiplyResult)
	fmt.Println(powResult)

}
