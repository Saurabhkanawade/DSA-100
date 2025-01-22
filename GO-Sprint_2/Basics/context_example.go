package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, 4, 5, 6}

	newSlice := append(ints[0:2], ints[3:]...)

	fmt.Println(newSlice)

}
