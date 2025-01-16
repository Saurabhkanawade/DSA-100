package main

import "fmt"

//pipeline example with goroutine / concurrency
func numsToChannel(num []int) <-chan int {
	newChan := make(chan int)

	go func() {
		for _, value := range num {
			newChan <- value
		}
		close(newChan)
	}()
	return newChan
}

func squareNums(n <-chan int) <-chan int {
	newChan := make(chan int)

	go func() {
		for value := range n {
			result := value * value
			newChan <- result
		}
		close(newChan)
	}()
	return newChan
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//stage1 add nums to channel
	numChan := numsToChannel(nums)

	//stage 2 square the nums
	squareChan := squareNums(numChan)

	//state 3 to print the results
	for result := range squareChan {
		fmt.Println(result)
	}
}
