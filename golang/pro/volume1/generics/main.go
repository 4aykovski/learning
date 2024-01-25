package main

import "fmt"

func myPrint[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
	fmt.Println()
}

func main() {
	ints := []int{1, 2, 3, 4}
	strs := []string{"a", "b", "c"}

	myPrint(ints)
	myPrint(strs)
}
