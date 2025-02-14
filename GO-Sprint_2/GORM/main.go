package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var p interface{}
	p = Person{name: "Alice", age: 25}

	// Type assertion to access the name field
	fmt.Println(p.(Person))

	// Type assertion to access the age field
	fmt.Println(p.(Person).age)
}
