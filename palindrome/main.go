package main

import "fmt"

func palindromicSubstrings(s string) []string {
	// Edge case: empty string
	var result []string

	n := len(s)

	// Helper function to expand around center
	expandAroundCenter := func(left, right int) {
		// Expand around center: left and right pointers
		// Move left pointer to the left and right pointer to the right
		// Check if the characters at left and right pointers are equal
		// If equal, add the substring to the result and continue expanding
		// If not equal, break the loop
		for left >= 0 && right < n && s[left] == s[right] {
			substr := s[left : right+1]
			if len(substr) > 1 { // Ignore single character substrings
				result = append(result, substr)
			}
			left--
			right++
		}
	}

	// Iterate through each character as potential center
	for i := 0; i < n; i++ {
		// Odd length palindromes (center at s[i])
		expandAroundCenter(i, i)
		// Even length palindromes (center between s[i] and s[i+1])
		expandAroundCenter(i, i+1)
	}

	return result
}

func main() {
	s := "babad"
	fmt.Println("Input:", s)
	fmt.Println("Palindromic substrings (excluding single characters):", palindromicSubstrings(s))
}
