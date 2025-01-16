package main

import "fmt"

func main() {

	myChannel := make(chan string)

	go func() {
		defer close(myChannel)

		fmt.Println("putting the data in channel")
		myChannel <- "data"

	}()

	channelData := <-myChannel

	fmt.Println(channelData)
	fmt.Println("main function")
}
