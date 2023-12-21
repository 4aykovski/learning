package main

import "errors"

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

func main() {

}
