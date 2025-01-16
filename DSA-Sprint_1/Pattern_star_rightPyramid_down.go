package main

import "fmt"

func PyramidStarPattern() {
	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	PyramidStarPattern()
}
