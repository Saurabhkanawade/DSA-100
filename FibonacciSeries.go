package main

import "fmt"

func PrintFibonacciSeries(limit int) []int {

	if limit < 1 {
		return []int{}
	}

	fibonacciSeries := []int{0, 1}

	for i := range limit {
		next := fibonacciSeries[i] + fibonacciSeries[i+1]
		fibonacciSeries = append(fibonacciSeries, next)
	}
	return fibonacciSeries
}

func main() {
	fmt.Println(PrintFibonacciSeries(10))
}
