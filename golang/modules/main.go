package main

import "modules/storage"

func spawnEmployees(s storage.Storage) {
	for i := 0; i < 10; i++ {
		err := s.Insert(storage.Employee{Id: i})
		if err != nil {
			return
		}
	}
}

func main() {
	var s1 storage.Storage
	var s2 storage.Storage
	s1 = storage.NewMemoryStorage()
	s2 = storage.NewDumbStorage()
	spawnEmployees(s1)
	spawnEmployees(s2)
}
