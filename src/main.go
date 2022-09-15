package main

import (
	"fmt"
)

// interface that hold the functions that my structs are sharing
type area2D interface {
	area() float64
}

// create structs
type square struct {
	base float64
}

type rectangle struct {
	base float64
	high float64
}

// function to calculate de area
func (s square) area() float64 {
	return s.base * s.base
}

func (r rectangle) area() float64 {
	return r.base * r.high
}

// function to calculate de area and instantiate the interface
func calculate(a area2D) {
	fmt.Println("Area: ", a.area())
}

func main() {
	//instantiate the structs

	var mySquare square
	var myRectangle rectangle

	mySquare.base = 2
	myRectangle.base = 2
	myRectangle.high = 4

	calculate(mySquare)
	calculate(myRectangle)

	//Interface list
	myInterfaceList := []interface{}{"Hola", 1, 12, true}

	fmt.Println(myInterfaceList...)
}
