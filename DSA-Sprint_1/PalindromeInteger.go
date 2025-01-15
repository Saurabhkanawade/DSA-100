package main

import (
	"fmt"
)

func PalindromeInteger(n int) string {
	orignalNum := n
	var reverseInt = 0

	if n < 0 {
		return fmt.Sprintf("The given string %s is NOT Palindrome", n)

	}

	for n > 0 {
		lastDigit := n % 10
		reverseInt = reverseInt*10 + lastDigit
		n = n / 10
	}

	if reverseInt == orignalNum {
		return fmt.Sprintf("The given No %d is Palindrome", orignalNum)
	}
	return fmt.Sprintf("The given No %d is NOT Palindrome", orignalNum)
}

func main() {
	fmt.Println(PalindromeInteger(8920298))
}
