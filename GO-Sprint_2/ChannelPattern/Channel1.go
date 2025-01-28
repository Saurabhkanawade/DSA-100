package main

import "fmt"

// channel without the waitgroup
func main() {

	myChannel := make(chan string, 1)

	go func() {
		defer close(myChannel)
		fmt.Println("putting the data in channel")
		myChannel <- "data"
		myChannel <- "data"
	}()

	channelData := <-myChannel

	fmt.Println(channelData)
	fmt.Println(channelData)
	fmt.Println("main function")
}
