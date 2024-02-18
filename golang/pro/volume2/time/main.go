package main

import (
	"fmt"
	"time"
)

func test(test int) {

}

func main() {
	myTime, _ := time.Parse("02 January 2006", "22 September 2004")
	fmt.Printf("%v", myTime.String())
}
