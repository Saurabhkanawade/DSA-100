package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := [5]int{31, 321, 43, 24, 45}
	numsSlice := nums[:]

	slices.SortFunc(numsSlice, func(a, b int) int {
		return a - b
	})

	fmt.Println(numsSlice)
}
