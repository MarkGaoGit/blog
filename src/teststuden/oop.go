package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {

	r1 := Rectangle{2, 4}
	rarea := r1.area()

	c1 := Circle{5}
	carea := c1.area()

	fmt.Println("Rectangle is Area :", rarea)
	fmt.Println("Circle is Area :", carea)

}
