package main

import "fmt"

type developer struct {
	name string
}

func (d developer) greeting() {
	fmt.Println("My name is " + d.name)
}

type employee struct {
	name string
}

func (e employee) greeting() {
	fmt.Printf("My name is %s and i'm a employee", e.name)
}

type person interface {
	greeting()
}

func massGreeting(persons []person) {
	for _, human := range persons {
		human.greeting()
	}
}

func main() {
	myDev := developer{name: "egor"}
	myEmp := employee{name: "pasha"}
	humans := []person{myDev, myEmp}
	massGreeting(humans)
}
