package main

import "fmt"

var roman = make(map[byte]int)

func init() {
	roman['I'] = 1
	roman['V'] = 5
	roman['X'] = 10
	roman['L'] = 50
	roman['C'] = 100
	roman['D'] = 500
	roman['M'] = 1000
}

func romanToInt(s string) int {

	total := 0

	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && roman[s[i]] < roman[s[i+1]] {
			total = total - roman[s[i]]
		} else {
			total = total + roman[s[i]]
		}
	}

	return total

}

func main() {
	fmt.Println("Roman to Integer conversion")

	// Test cases
	fmt.Println("III:", romanToInt("III"))     // Output: 3
	fmt.Println("IV:", romanToInt("IV"))       // Output: 4
	fmt.Println("IX:", romanToInt("IX"))       // Output: 9
	fmt.Println("LVIII:", romanToInt("LVIII")) // Output: 58
	fmt.Println("MCMXCIV:", romanToInt("MCMXCIV"))

}
