package main

import "fmt"

// With Liskov substitution principle

type Bird interface {
	MakeSound() string
}

type FlyingBird interface {
	Bird
	Fly() string
}

type Pigeon struct {
}

func (p *Pigeon) Fly() string {
	return "Pigeon can fly"
}

func (p *Pigeon) MakeSound() string {
	return "coo"
}

type Penguin struct {
}

func (p *Penguin) MakeSound() string {
	return "Penguin can make sound"
}

func ProcessFlyingBird(flyingBird FlyingBird) string {
	return flyingBird.Fly()
}

func ProcessBird(bird Bird) string {
	return bird.MakeSound()
}

func main() {

	birds := []Bird{
		&Penguin{},
		&Pigeon{},
	}

	for _, bird := range birds {
		fmt.Println(ProcessBird(bird))
	}

	birdFlying := []FlyingBird{
		&Pigeon{},
	}

	for _, bird := range birdFlying {
		fmt.Println(ProcessFlyingBird(bird))
	}
}

//without LSP
//
//type Bird interface {
//	Fly() string
//}
//
//type Pigeon struct{}
//
//func (p *Pigeon) Fly() string {
//	return "Pigeon can fly"
//}
//
//type Penguin struct{}
//
//func (p *Penguin) Fly() string {
//	return "Penguin can fly"
//}
