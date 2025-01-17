package main

import (
	"fmt"
	"strconv"
)

func PalindromeString1(n int) {
	str := strconv.Itoa(n)
	reverseInt := ""
	for _, ch := range str {
		reverseInt = string(ch) + reverseInt
	}
	reversedInt, _ := strconv.Atoi(reverseInt)
	if reversedInt == n {
		fmt.Println("The given no is Palindrome")
	} else {
		fmt.Println("The given no is NOT Palindrome")
	}
}

func main() {
	PalindromeString1(515)
}
