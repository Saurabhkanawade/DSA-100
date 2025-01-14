package main

import "fmt"

func evenOdd(n int) string {
	var answer string

	if n%2 == 0 {
		answer = fmt.Sprintf("The given no %d is Even", n)
	} else {
		answer = fmt.Sprintf("The given no %d is Odd", n)
	}
	return answer
}

func main() {
	fmt.Println(evenOdd(5))
}
