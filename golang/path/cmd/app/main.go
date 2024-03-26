package main

import "fmt"

func calc(i *int) {
	*i = *i - 1*10
}

func main() {
	a := 1
	b := &a
	*b++
	c := &b
	*b++
	_ = c
	fmt.Println(a)
}
