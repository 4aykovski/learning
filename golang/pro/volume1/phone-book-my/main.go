package main

import (
	"flag"
	"fmt"
)

type entry struct {
	name     string
	surname  string
	phoneNum string
}

var data []entry

func search(key string) *entry {
	for i, v := range data {
		if v.surname == key {
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}
func main() {
	listFlag := flag.Bool(
		"list",
		false,
		"Use this flag to print all data in phone book",
	)
	searchFlag := flag.String(
		"search",
		"",
		"Use this flag to search a phone in phone book depending on surname",
	)
	flag.Parse()

	if *listFlag && *searchFlag != "" {
		panic("Was given two flags at once!")
	}

	if !*listFlag && *searchFlag == "" {
		panic("Wan't given any flags!")
	}

	data = append(data, entry{"Mihalis", "Tsoukalos", "2109416471"})
	data = append(data, entry{"Mary", "Doe", "2109416871"})
	data = append(data, entry{"John", "Black", "2109416123"})

	switch {
	case *searchFlag != "":
		result := search(*searchFlag)
		if result == nil {
			fmt.Println("Not found")
			return
		}
		fmt.Println(result)
	case *listFlag:
		list()
	}
}
