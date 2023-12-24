package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int = 0

func main() {
	//t := time.Now()
	//
	//fmt.Printf("Start: %v\n", t)
	//
	//fmt.Println("Processing...")
	//
	//result1 := make(chan int)
	//result2 := make(chan int)
	//
	//go calculateSomething(1000, result1)
	//go calculateSomething(2000, result2)
	//
	//fmt.Printf("Result1: %v\n", <-result1)
	//fmt.Printf("Result2: %v\n", <-result2)
	//
	//fmt.Printf("End: %v\n", time.Since(t))

	//numbers := make(chan int)
	//
	//go generateNumbers(1000, numbers)
	//
	//for n := range numbers {
	//	fmt.Println(n)
	//}

	//numbers := make(chan int)
	//
	//select {
	//case n := <-numbers:
	//	fmt.Println(n)
	//default:
	//	fmt.Println("Канал пуст.")
	//}

	//result := make(chan uint64)
	//
	//go func() {
	//	for _, r := range "-\\|/" {
	//		fmt.Println(string(r))
	//		time.Sleep(time.Millisecond * 100)
	//	}
	//}()
	//
	//go factorial(65, result)
	//
	//fmt.Println(<-result)

	//channel := make(chan int)
	//
	//close(channel)
	//
	//channel <- 5

	fmt.Println("Start...")
	var wg sync.WaitGroup
	wg.Add(5)
	var mutex sync.Mutex
	for i := 0; i < 5; i++ {
		go work(i, &mutex, &wg)
	}

	wg.Wait()
	fmt.Println("End...")
}

func work(i int, m *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Lock()
	counter = 0
	for k := 0; k < 5; k++ {
		counter++
		fmt.Printf("Number - %d, counter - %d\n", i, counter)
	}
	m.Unlock()
}

func factorialRec(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	return n * factorialRec(n-1)
}

func factorial(n uint64, res chan uint64) {
	fact := factorialRec(n)
	res <- fact
}

func generateNumbers(n int, nums chan int) {
	for i := 0; i < n; i++ {
		fmt.Println("process")
		nums <- i * 2
	}
	fmt.Println("end")
	close(nums)
}

func calculateSomething(n int, res chan int) {
	t := time.Now()

	result := 0
	for i := 0; i < n; i++ {
		result += i * 2
		time.Sleep(time.Millisecond * 3)
	}

	fmt.Printf("Time: %v\n", time.Since(t))
	res <- result
}
