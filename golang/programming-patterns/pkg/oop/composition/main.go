package main

import "fmt"

type car struct {
	engine
	wheels [4]wheel
}

func newCar() *car {
	engine := engine{}
	wheels := [4]wheel{wheel{}, wheel{}, wheel{}, wheel{}}
	return &car{
		engine: engine,
		wheels: wheels,
	}
}

func (c car) drive() {
	c.engine.drive()
	for _, wheel := range c.wheels {
		wheel.drive()
	}
}

type wheel struct{}

func (w wheel) drive() {
	fmt.Println("Wheel is wheeling")
}

type engine struct{}

func (e engine) drive() {
	fmt.Println("Engine is working")
}

func main() {
	car := newCar()
	car.drive()
}
