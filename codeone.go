//Solving the below problem in go.
// https://www.hackerrank.com/challenges/validating-credit-card-number/problem

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isValidCreditCardNumber(number string) bool {

	// Handling special cases that are separated by one hyphen "-".
	parts := strings.Split(number, "-")

	if len(parts) > 1 {
		if len(parts) != 4 {
			return false
		}
		for _, part := range parts {
			if len(part) != 4 {
				return false
			}
		}
		number = strings.Join(parts, "")
	}

	// Checking if number contains exactly 16 digits
	if len(number) != 16 {
		return false
	}

	// Checking if number starts with 4, 5, or 6
	if number[0] != '4' && number[0] != '5' && number[0] != '6' {
		return false
	}

	// Checking for non-digit characters
	for _, ch := range number {
		if ch < '0' || ch > '9' {
			return false
		}
	}

	// Checking for 4 plus repeated digits
	for i := 0; i <= 12; i++ {
		if number[i] == number[i+1] && number[i+1] == number[i+2] && number[i+2] == number[i+3] {
			return false
		}
	}

	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter number of credit cards to validate:")

	scanner.Scan()

	n := 0

	fmt.Sscanf(scanner.Text(), "%d", &n)

	for i := 0; i < n; i++ {

		fmt.Println("Enter credit card number:")

		scanner.Scan()
		cardNumber := scanner.Text()

		if isValidCreditCardNumber(cardNumber) {

			fmt.Println("Valid")

		} else {

			fmt.Println("Invalid")
		}
	}
}
