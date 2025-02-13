package main

import "fmt"

func oddEven(odd, even chan int, n int) {
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(odd)
	close(even)
}

func main() {
	even := make(chan int)
	odd := make(chan int)

	go oddEven(odd, even, 10)

	for range 10 {
		select {
		case value, ok := <-even:
			if ok {
				fmt.Println("even", value)
			}
		case value, ok := <-odd:
			if ok {
				fmt.Println("odd", value)
			}
		}
	}
}
