package main

import "fmt"

func SumInteger(ints []int) {
	sum := 0

	for _, value := range ints {
		sum += value
	}
	fmt.Println(sum)
}

func main() {
	SumInteger([]int{1, 2, 3, 4})
}
