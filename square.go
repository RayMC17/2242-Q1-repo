package main

import (
"log"
)

/*Write a Go function named 'square' that accepts a parameter of type float64
   that represents the side of the square. The 'square' function should return
   the area and the perimeter of the square when called.*/

func Square(side float64)(float64,float64) {
area:= (side * side) 
perimeter:= 4 * side
return area, perimeter
}

func main(){
	log.Println(Square(7))
}
