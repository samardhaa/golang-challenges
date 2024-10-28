// Program to reverse the input of the given string

package main

import "fmt"

func main() {
	fmt.Println("Program to reverse the given string")

	input := "hello there!"

	fmt.Println("given input:", input)

	fmt.Println(reverseString(input))

}

func reverseString(s string) string {

	result := []rune(s)

	i := 0
	j := len(result) - 1

	for i < j {
		temp := result[i]
		result[i] = result[j]
		result[j] = temp

		i++
		j--
	}

	return string(result)

}

// if your goal is character-based manipulation, []rune is the right choice. But for byte-level manipulation or ASCII-only data, []byte works well
