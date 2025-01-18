package main

import "fmt"

type Number interface {
	int | int32 | int64 | float64 | float32
}

func Test[T Number](n T) {
	fmt.Println("The number is :", n)
}

func main() {
	Test(2.4)
}
