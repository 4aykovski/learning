package main

import (
	"fmt"
	"shapes/shapes"
)

func main() {
	rectangle := shapes.NewRectangle(10, 20)
	triangle := shapes.NewTriangle(10, 20, 20)

	areas, err := shapes.GetAllAreas(rectangle, triangle)
	if err != nil {
		panic(err)
	}

	for _, area := range areas {
		fmt.Printf("Area: %f64\n", area)
	}
}
