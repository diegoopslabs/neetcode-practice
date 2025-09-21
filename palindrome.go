package main

import (
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	s = normalizeString(s)
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func normalizeString(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			builder.WriteRune(unicode.ToLower(r))
		}
	}
	return builder.String()
}

func IsPalindromeRecursive(s string) bool {
	s = normalizeString(s)
	return isPalindromeHelper(s, 0, len(s)-1)
}

func isPalindromeHelper(s string, left, right int) bool {
	if left >= right {
		return true
	}
	if s[left] != s[right] {
		return false
	}
	return isPalindromeHelper(s, left+1, right-1)
}