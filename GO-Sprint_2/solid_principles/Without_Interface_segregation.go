package main

import "fmt"

type CarWash interface {
	Wash() string
	Repair() string
	Modify() string
}

type CarDetails struct {
	name string
}

func (c *CarDetails) Repair() string {
	//TODO implement me
	panic("implement me")
}

func (c *CarDetails) Modify() string {
	//TODO implement me
	panic("implement me")
}

func NewCar(name string) CarWash {
	return &CarDetails{
		name: name,
	}
}

func (c *CarDetails) Wash() string {
	return "car is washing"
}

func main() {

	car := NewCar("Maruti")
	fmt.Println(car.Wash())

	s := CarDetails{
		name: "dk",
	}
	fmt.Println(s.Wash())

}
