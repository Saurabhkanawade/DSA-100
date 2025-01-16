package main

import "fmt"

func StarPyramid() {
	var n = 10
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func StarPyramid2() {
	var n = 10
	for i := 0; i < n; i++ {
		for j := n; j > i; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	StarPyramid()
	StarPyramid2()
}
