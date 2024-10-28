// Problem Name: Find Longest Palindromic Substring
// Problem Statement: Given a string s, find and return the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
// Example(s): the longest palindrome in the string "ababd" is "aba" or "bab" (either is acceptable) and we would want either returned
// Example Inputs: asdfsda, hellollehasdf, abahellollehaba, abacdcf

package main

import "fmt"

func main() {
	fmt.Println("Longest Palindromic substring")
	fmt.Println("asdfsda")
	fmt.Println("asdfsda")
	fmt.Println("abahellollehaba")
	fmt.Println("abacdcf")

	inputs := []string{
		"asdfsda",
		"abahellollehaba",
		"abacdcf",
	}

	for _, input := range inputs {
		result := longestPalindromicSubstring(input)
		fmt.Printf("Input: %s  ---- %s\n", input, result)
	}
}

func longestPalindromicSubstring(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	dp := make([][]bool, n)

	for i := range dp {
		dp[i] = make([]bool, n)
	}

	start := 0
	maxLen := 1

	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			maxLen = 2
		}
	}

	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1

			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				start = i
				maxLen = length
			}
		}
	}

	return s[start : start+maxLen]

}
