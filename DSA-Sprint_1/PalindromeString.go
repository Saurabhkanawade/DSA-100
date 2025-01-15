package main

import "fmt"

func PalindromeString(str string) string {
	reverseString := ""

	for _, ch := range str {
		reverseString = string(ch) + reverseString
	}
	if reverseString == str {
		return fmt.Sprintf("The given string %s is Palindrome", str)
	}
	return fmt.Sprintf("The given string %s is NOT Palindrome", str)
}

func main() {
	fmt.Println(PalindromeString("radar"))
}
