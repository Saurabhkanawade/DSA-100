package main

import "fmt"

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func main() {
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		return
	}
	fmt.Printf("The factorial of the %d is : %d", number, Factorial(number))
}
