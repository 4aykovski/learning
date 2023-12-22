package main

import (
	"errors"
	"fmt"
)

type employee struct {
	id     int
	name   string
	age    int
	salary int
}

type storage interface {
	insert(e employee) error
	get(id int) (employee, error)
	delete(id int) error
}

type dumbStorage struct {
}

func newDumbStorage() *dumbStorage {
	return &dumbStorage{}
}

func (d *dumbStorage) insert(e employee) error {
	fmt.Println("Employee was inserted")
	return nil
}

func (d *dumbStorage) get(id int) (employee, error) {
	e := employee{
		id: id,
	}
	return e, nil
}

func (d *dumbStorage) delete(id int) error {
	fmt.Println("Employee was deleted")
	return nil
}

type memoryStorage struct {
	data map[int]employee
}

func newMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]employee),
	}
}

func (m *memoryStorage) insert(e employee) error {
	m.data[e.id] = e

	return nil
}

func (m *memoryStorage) get(id int) (employee, error) {
	e, exists := m.data[id]
	if !exists {
		return employee{}, errors.New("employee not found")
	}
	return e, nil
}

func (m *memoryStorage) delete(id int) error {
	delete(m.data, id)
	return nil
}

func spawnEmployees(s storage) {
	for i := 0; i < 10; i++ {
		err := s.insert(employee{id: i})
		if err != nil {
			return
		}
	}
}

func main() {
	var s1 storage
	var s2 storage
	s1 = newMemoryStorage()
	s2 = newDumbStorage()
	spawnEmployees(s1)
	spawnEmployees(s2)
}
