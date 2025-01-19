package main

import "fmt"

type Car struct {
	Brand      string
	Seats      int
	Engine     string
	Sunroof    bool
	Gps        bool
	HeatedSeat bool
}

type CarBuilder struct {
	brand      string
	seats      int
	engine     string
	sunroof    bool
	gps        bool
	heatedSeat bool
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (c *CarBuilder) setBrand(brand string) *CarBuilder {
	c.brand = brand
	return c
}

func (c *CarBuilder) addSeats(seat int) *CarBuilder {
	c.seats = seat
	return c
}

func (c *CarBuilder) setEngine(engine string) *CarBuilder {
	c.engine = engine
	return c
}

func (c *CarBuilder) addSunroof(sunroof bool) *CarBuilder {
	c.sunroof = sunroof
	return c
}

func (c *CarBuilder) addHeatedSeat(heatedSeat bool) *CarBuilder {
	c.heatedSeat = heatedSeat
	return c
}

func (c *CarBuilder) addGps(gps bool) *CarBuilder {
	c.gps = gps
	return c
}

func (c *CarBuilder) Build() *Car {
	return &Car{
		Brand:      c.brand,
		Seats:      c.seats,
		Engine:     c.engine,
		Sunroof:    c.sunroof,
		Gps:        c.gps,
		HeatedSeat: c.heatedSeat,
	}
}

func main() {
	builder := NewCarBuilder()

	car := builder.setBrand("Suzuki").
		addSeats(4).
		setEngine("4 Litter turbo").
		addSunroof(true).
		addGps(true).
		addHeatedSeat(true).
		Build()

	fmt.Printf("The car details %+v\n", car)
}
