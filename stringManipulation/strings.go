package main

import "strings"

// ReverseString reverses the input string.
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// InsertString inserts a substring into a string at a specified position.
func InsertString(s, insert string, index int) string {
	if index > len(s) {
		index = len(s)
	}
	return s[:index] + insert + s[index:]
}

// SearchString searches for the first occurrence of a substring in a string.
func SearchString(s, search string) int {
	return strings.Index(s, search)
}

// RemoveSubstring removes the first occurrence of a substring from a string.
func RemoveSubstring(s, remove string) string {
	index := SearchString(s, remove)
	if index == -1 {
		return s
	}
	return s[:index] + s[index+len(remove):]
}
