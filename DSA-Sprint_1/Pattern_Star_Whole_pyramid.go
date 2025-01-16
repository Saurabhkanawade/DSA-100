package main

import "fmt"

func WholePyramid() {

	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	WholePyramid()
}
