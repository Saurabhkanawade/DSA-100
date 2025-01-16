package main

import "fmt"

func StartLeftPyramid() {

	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	StartLeftPyramid()
}
