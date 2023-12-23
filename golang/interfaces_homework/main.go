package main

import (
	"errors"
	"fmt"
	"math"
)

type GeometricalShape interface {
	Area() (float64, error)
	Perimeter() (float64, error)
}

type Rectangle struct {
	width  float64
	height float64
}

func NewRectangle(width float64, height float64) *Rectangle {
	return &Rectangle{width: width, height: height}
}

func (r *Rectangle) Area() (float64, error) {
	if r.width <= 0 || r.height <= 0 {
		return 0, errors.New("Invalid sizes")
	}
	return r.width * r.height, nil
}

func (r *Rectangle) Perimeter() (float64, error) {
	if r.width <= 0 || r.height <= 0 {
		return 0, errors.New("Invalid sizes")
	}
	return (r.width + r.height) * 2, nil
}

type Triangle struct {
	firstSide  float64
	secondSide float64
	thirdSide  float64
}

func NewTriangle(firstSide float64, secondSide float64, thirdSide float64) *Triangle {
	return &Triangle{firstSide: firstSide, secondSide: secondSide, thirdSide: thirdSide}
}

func (t *Triangle) Area() (float64, error) {
	if t.firstSide <= 0 || t.secondSide <= 0 || t.thirdSide <= 0 {
		return 0, errors.New("Invalid sizes")
	}
	perimeter, err := t.Perimeter()
	if err != nil {
		return 0, errors.New("Invalid sizes")
	}
	halfPerimeter := perimeter / 2
	return math.Sqrt(halfPerimeter * (halfPerimeter - t.firstSide) * (halfPerimeter - t.secondSide) * (halfPerimeter - t.thirdSide)), nil
}

func (t *Triangle) Perimeter() (float64, error) {
	if t.firstSide <= 0 || t.secondSide <= 0 || t.thirdSide <= 0 {
		return 0, errors.New("Invalid sizes")
	}
	return t.firstSide + t.secondSide + t.thirdSide, nil
}

func getAllAreas(shapes ...GeometricalShape) ([]float64, error) {
	var areas []float64
	for _, shape := range shapes {
		area, err := shape.Area()
		if err != nil {
			return nil, err
		}
		areas = append(areas, area)
	}
	return areas, nil
}

func main() {
	rectangle := NewRectangle(10, 20)
	triangle := NewTriangle(2, 2, 1)

	areas, err := getAllAreas(rectangle, triangle)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	for _, area := range areas {
		fmt.Printf("Area: %v\n", area)
	}
}
