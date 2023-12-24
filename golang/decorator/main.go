package main

import "fmt"

func main() {
	summa := decorator(summa)
	sum := summa(1, 2, 3, 4, 5)
	fmt.Println(sum)
	fmt.Println([]int{1, 2, 3, 4})
}

func summa(n ...int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}

func decorator(f func(...int) int) func(...int) int {
	return func(args ...int) int {
		fmt.Println("Start function...")
		res := f(args...)
		fmt.Println("End function...")
		return res
	}
}
