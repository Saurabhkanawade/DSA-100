package main

import (
	"fmt"
	"sync"
)

func printOdd(chOdd, chEven chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		<-chOdd // Wait for odd signal
		fmt.Printf("odd func : %d\n", i)
		chEven <- true // Signal even channel
	}
}

func printEven(chEven, chOdd chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		<-chEven // Wait for even signal
		fmt.Printf("even func : %d\n", i)
		if i == 10 {
			close(chOdd)
			close(chEven)
			return
		}
		chOdd <- true // Signal odd channel
	}
}

func main() {
	chOdd := make(chan bool)
	chEven := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(2)

	go printOdd(chOdd, chEven, &wg)
	go printEven(chEven, chOdd, &wg)

	// Start the process by signaling the odd channel
	chOdd <- true

	// Wait for Goroutines to finish
	wg.Wait()
}
