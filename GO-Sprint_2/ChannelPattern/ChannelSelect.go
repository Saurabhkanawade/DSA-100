package main

import "fmt"

func main() {
	myChan1 := make(chan string)
	myChan2 := make(chan string)

	go func() {
		myChan1 <- "this is chan 1"
	}()

	go func() {
		myChan2 <- "this is chan 2"
	}()

	//if the both values are ready in channel then select statement randomly print.
	select {
	case myChan1Data := <-myChan1:
		fmt.Println(myChan1Data)
	case myChan2Data := <-myChan2:
		fmt.Println(myChan2Data)

	}

}
