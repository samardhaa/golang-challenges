package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("string reverse function")

	fmt.Print("Enter the string: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	fmt.Println("Given string :", input)

	fmt.Println("Reverse of the given string :", stringReverse(input))

}

func stringReverse(s string) string {

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
