package main

import "fmt"

func main() {
	charChannel := make(chan string, 4)
	chars := []string{"a", "b", "c", "d"}

	for _, ch := range chars {
		select {
		case charChannel <- ch:
		}
	}
	close(charChannel)

	for value := range charChannel {
		fmt.Println(value)
	}

}
