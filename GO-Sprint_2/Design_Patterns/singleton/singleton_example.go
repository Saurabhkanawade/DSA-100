package main

import (
	"fmt"
	"sync"
)

var mt = &sync.Mutex{}

type car struct {
	name string
}

var singleCar *car

func getSingleCar() *car {
	if singleCar == nil {
		mt.Lock()
		defer mt.Unlock()
		singleCar = &car{name: "swift"}
		fmt.Println("created the single instance.", singleCar)
	} else {
		fmt.Println("instance is already present.", singleCar)
	}
	return singleCar
}

func main() {
	for i := 0; i < 30; i++ {
		getSingleCar()
	}
}
