package storage

import (
	"errors"
	"fmt"
)

type Employee struct {
	Id     int
	Name   string
	Age    int
	Salary int
}

type Storage interface {
	Insert(e Employee) error
	Get(id int) (Employee, error)
	Delete(id int) error
}

type DumbStorage struct {
}

func NewDumbStorage() *DumbStorage {
	return &DumbStorage{}
}

func (d *DumbStorage) Insert(e Employee) error {
	fmt.Println("Employee was inserted")
	return nil
}

func (d *DumbStorage) Get(id int) (Employee, error) {
	e := Employee{
		Id: id,
	}
	return e, nil
}

func (d *DumbStorage) Delete(id int) error {
	fmt.Println("Employee was deleted")
	return nil
}

type MemoryStorage struct {
	data map[int]Employee
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[int]Employee),
	}
}

func (m *MemoryStorage) Insert(e Employee) error {
	m.data[e.Id] = e

	return nil
}

func (m *MemoryStorage) Get(id int) (Employee, error) {
	e, exists := m.data[id]
	if !exists {
		return Employee{}, errors.New("Employee not found")
	}
	return e, nil
}

func (m *MemoryStorage) Delete(id int) error {
	delete(m.data, id)
	return nil
}
