package main

import(
	"log"
	"math"
)

/*Num 6: Create a struct type named 'circle' that has one field named 'radius' of type float64.*/

type circle struct{
	radius float64
}

/*Num 7: Create a method on type 'circle' named 'area' that calculates and
   returns the area of a circle.*/

   func (c circle)area()float64 {
	formula:= math.Pi * (float64(c.radius) * float64(c.radius))
	return formula
   }

/*Num 8: Create a method on type 'circle' named 'perimeter' that calculates and
   returns the area of a circle.*/

   func (c circle)perimeter()float64{
	formula:= 2* (22/7.0)* c.radius
	return formula
   }
//created a variable of type circle
   func main() {
	shape_b := circle{
	radius: 9,
	}
	
//called area and perimeter
	log.Println(shape_b.area())
	log.Println(shape_b.perimeter())

}
