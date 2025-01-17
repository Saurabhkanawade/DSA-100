package main

import "fmt"

func FactorialNumber(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	fmt.Print(FactorialNumber(5))
}
