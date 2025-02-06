package main

import "fmt"

type Operation interface {
	Calculator(a, b int)
}

type Addition struct {
}

func (Addition) Calculator(a, b int) {
	fmt.Println("The addition of the values is :", a+b)
}

type Subtraction struct {
}

func (Subtraction) Calculator(a, b int) {
	fmt.Println("The Subtraction of the values is :", a-b)
}

type Multiplication struct {
}

func (Multiplication) Calculator(a, b int) {
	fmt.Println("The Multiplication of the values is :", a*b)
}

type Divide struct {
}

func (Divide) Calculator(a, b int) {
	fmt.Println("The Divide of the values is :", a/b)
}

func main() {

	op := []Operation{
		Addition{},
		Subtraction{},
		Multiplication{},
		Divide{},
	}

	for _, op := range op {
		op.Calculator(5, 2)
	}
}
