package main

import (
	"fmt"
	"sync"
)

func PrintOddNumbers(oddChan chan bool, evenChan chan bool, result chan int, s *sync.WaitGroup, n int) {
	defer s.Done()

	for i := 1; i <= 10; i = i + 2 {
		<-oddChan
		result <- i
		evenChan <- true
	}
}

func PrintEvenNumbers(evenChan chan bool, oddChan chan bool, result chan int, s *sync.WaitGroup, n int) {
	defer s.Done()
	for i := 2; i <= n; i = i + 2 {
		<-evenChan
		result <- i
		if i == 10 {
			close(evenChan)
			close(oddChan)
			close(result)
			return
		}
		oddChan <- true
	}

}

func main() {

	var wg sync.WaitGroup

	iteration := 10
	oddChan := make(chan bool)
	evenChan := make(chan bool)
	sequentialData := make(chan int)

	wg.Add(2)
	go PrintOddNumbers(oddChan, evenChan, sequentialData, &wg, iteration)
	go PrintEvenNumbers(evenChan, oddChan, sequentialData, &wg, iteration)

	oddChan <- true

	for value := range sequentialData {
		fmt.Println(value)
	}

	wg.Wait()
}
