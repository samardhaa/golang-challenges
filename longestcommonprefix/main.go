package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"hello", "hell", "helieum", "helloworld"}

	fmt.Println("string matching is", stringMatch(s))

}

func stringMatch(s []string) string {

	prefix := s[0]

	if prefix == "" {
		return "empty string"
	}

	for i := 1; i < len(s); i++ {
		if !strings.HasPrefix(s[i], prefix) {
			prefix = prefix[:len(prefix)-1]
		} else if prefix == "" {
			return "empty string"
		}

	}

	return prefix

}
