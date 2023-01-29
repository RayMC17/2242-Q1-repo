package main

import (
	"log"
	"math"
)

/*Num 1: Create a struct type named 'triangle' that has two fields named
base and height both of type float64.*/
type triangle struct {
	base   float64
	height float64
}

/*Num 2: Create a method on type 'triangle' named 'area' that calculates and
returns the area of a triangle.*/

func (t triangle) area() float64 {
	return .5 * t.base * t.height
}

/*Num 3: Create a method on type 'triangle' named 'perimeter' that calculates and returns 
the perimeter of a triangle.*/

func (t triangle) perimeter() float64 {
	Val2 := math.Sqrt(math.Pow(t.base, 2) + math.Pow(t.height, 2))
	return Val2

}

/*Num 4: Inside the 'main' function create a variable of type 'triangle' */

func main() {
	shape := triangle{
		base:   10.5,
		height: 6.6,
	}
	
	//Num 5: Call the 'area' and 'perimeter' methods on the 'triangle' that you created.
	
	log.Println("Area =", shape.area())

	log.Println("Perimeter =", shape.perimeter())

}


