package main

import "fmt"

func sayHello(name string) string {
	return "Hello " + name + ",how are you doing ?"
}

func greetings(name string, greetings func(string2 string) string) {
	str := greetings(name)
	fmt.Println(str)
}

func main() {
	greetings("Saurabh", sayHello)
}
