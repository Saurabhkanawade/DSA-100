package main

import "fmt"

func Flag() bool {
	return 2%2 == 8
}

func main() {
	fmt.Println(Flag())
}
