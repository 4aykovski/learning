package main

import "fmt"

type company struct {
	name   string
	sphere string
}

func newCompany(name, sphere string) (*company, error) {
	if name == "" || sphere == "" {
		return nil, fmt.Errorf("invalid arguments")
	}
	var tmpCompany company
	tmpCompany.name = name
	tmpCompany.sphere = sphere
	return &tmpCompany, nil
}

func (c *company) setCompanyName(name string) {
	c.name = name
}

func (c *company) setSphere(sphere string) {
	c.sphere = sphere
}

type human struct {
	name       string
	age        int
	profession string
}

func newHuman(name, profession string, age int) (*human, error) {
	if name == "" || profession == "" || age <= 0 || age > 100 {
		return nil, fmt.Errorf("invalid arguments")
	}
	var tmpHuman human
	tmpHuman.name = name
	tmpHuman.age = age
	tmpHuman.profession = profession
	return &tmpHuman, nil
}

func (h *human) setName(name string) {
	h.name = name
}

func (h *human) setAge(age int) {
	h.age = age
}

func (h *human) setProfession(profession string) {
	h.profession = profession
}

func (h *human) printHumanInfo() {
	fmt.Printf("Name: %v\n", h.name)
	fmt.Printf("Age: %v\n", h.age)
	fmt.Printf("Proffesion: %v\n", h.profession)
}

type employee struct {
	human
	company
}

func newEmployee(human human, company company) *employee {
	var tmpEmployee employee
	tmpEmployee.human = human
	tmpEmployee.company = company
	return &tmpEmployee
}

func (e *employee) printAllInfo() {
	fmt.Printf("Name: %v\n", e.human.name)
	fmt.Printf("Age: %v\n", e.human.age)
	fmt.Printf("Proffesion: %v\n", e.profession)
	fmt.Printf("Company name: %v\n", e.company.name)
	fmt.Printf("Company sphere: %v\n", e.company.sphere)
}

func main() {
	human1, err := newHuman("Ali", "programmer", 12)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	company, err := newCompany("Google", "IT")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	employee := newEmployee(*human1, *company)
	employee.printHumanInfo()
	employee.printAllInfo()
}
