package main

import "fmt"

func main() {
	fmt.Println("Hello, Go Sandbox!")
	fmt.Println()

	testCases := []string{
		"racecar",
		"A man a plan a canal Panama",
		"race a car",
		"hello",
		"Madam",
		"Was it a car or a cat I saw?",
		"No 'x' in Nixon",
		"12321",
		"12345",
	}

	fmt.Println("Testing Palindrome Functions:")
	fmt.Println("=============================")

	for _, test := range testCases {
		iterative := IsPalindrome(test)
		recursive := IsPalindromeRecursive(test)

		fmt.Printf("%-30s -> Iterative: %-5t | Recursive: %-5t\n",
			fmt.Sprintf("\"%s\"", test), iterative, recursive)
	}
}