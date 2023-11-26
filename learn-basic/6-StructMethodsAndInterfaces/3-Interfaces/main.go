package main

import (
	"fmt"
)

// Declaring an Interface
// ######################
type shape interface {
	area() float64
	perimeter() float64
}

// Declaring a Square struct with its methods set
// ##############################################
type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (s Square) perimeter() float64 {
	return 4 * s.side
}
// #######################################



// Declaring a Rect struct with its methods set
// #############################################
type Rect struct {
	length, breadth float64
}

func (r Rect) area() float64 {
	return r.length * r.breadth
}

func (r Rect) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}
// #######################################

// Declaring a Circle struct with its methods set
// #############################################
const PI float64 = 3.14

type Circle struct {
	radius float64
}

func (r Circle) area() float64 {
	return  PI * r.radius * r.radius
}

func (r Circle) perimeter() float64 {
	return 2 * PI * r.radius
}
// #######################################



func printData(s shape) {
	fmt.Printf("s is %v is of type %T\n",s, s)
	fmt.Printf("s.area() = %.2f\n", s.area())
	fmt.Printf("s.perimeter() = %.2f\n", s.perimeter())
}

func main() {
	r := Rect{length: 3, breadth: 4}
	s := Square{side: 5}
	c := Circle{radius: 5}
	
	// Using the interface to print the data of a rectangle
	fmt.Println("// Using the interface to print the data of a rectangle")
	printData(r)
	fmt.Println()
	
	// Using the interface to print the data of a Square
	fmt.Println("// Using the interface to print the data of a Square")
	printData(s)
	fmt.Println()

	// Using the interface to print the data of a Circle
	fmt.Println("// Using the interface to print the data of a Circle")
	printData(c)
	fmt.Println()
}