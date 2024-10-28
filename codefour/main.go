// //"one+two-one-two-one+one-two"

// one = 1

// two = 2

// 1+2-1+2

package main

import (
	"fmt"
	"strconv"
)

func evaluateExpression(expression string) int {
	total := 0
	currentOp := "+"

	for i := 0; i < len(expression); i++ {
		char := expression[i]

		if char == '1' || char == '2' {
			value, _ := strconv.Atoi(string(char))

			if currentOp == "+" {
				total = total + value
			} else if currentOp == "-" {
				total = total - value
			}
		} else if char == '+' || char == '-' {
			currentOp = string(char)
		}
	}
	return total
}

func main() {
	expression := "1+1+2+2-2-1-2-1-1"

	result := evaluateExpression(expression)

	fmt.Printf("The result is %s is: %d\n", expression, result)

}
