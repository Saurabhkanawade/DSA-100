package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Single struct {
}

var singleInstance *Single

func getInstance() *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			fmt.Println("creating single instance now.")
			singleInstance = &Single{}
		} else {
			fmt.Println("single instance is already created.")
		}
	} else {
		fmt.Println("single instance is already created.")
	}
	return singleInstance
}

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}
}
