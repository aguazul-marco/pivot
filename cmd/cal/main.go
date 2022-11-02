package main

import (
	"fmt"
	"github.com/aguazul-marco/pivot/calculator"
)

func main() {

	addResult := calculator.Add(8, 8)
	subResult := calculator.Subtract(8, 8)
	multiplyResult := calculator.Multiply(8, 8)

	fmt.Println(addResult)
	fmt.Println(subResult)
	fmt.Println(multiplyResult)
}
