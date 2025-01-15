package main

import "fmt"

func isPrimeNumber(n int) string {
	var answer string
	var count int

	for num := range n {
		if n%(num+1) == 0 {
			count += 1
		}
	}

	if count == 2 {
		answer = fmt.Sprintf("The given no %d is Prime Number", n)
	} else {
		answer = fmt.Sprintf("The given no %d is Not a Prime Number", n)
	}

	return answer
}

func main() {
	fmt.Println(isPrimeNumber(11))
}
