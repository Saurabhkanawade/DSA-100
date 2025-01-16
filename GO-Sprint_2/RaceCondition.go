package main

import (
	"fmt"
	"sync"
)

var counter int

func IncreaseCounter(wg *sync.WaitGroup, mx *sync.Mutex) {
	defer wg.Done()
	defer mx.Unlock()
	mx.Lock()
	counter++
}

func main() {
	var wg sync.WaitGroup
	var mx sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go IncreaseCounter(&wg, &mx)
	}
	wg.Wait()
	fmt.Println(counter)
}
