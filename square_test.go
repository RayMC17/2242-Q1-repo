package main

import(
	"testing"
)

//testing Square Area
func TestSquare_Area(t *testing.T) {
    result, _ := Square(7.0)
    got := result
    want := 49.0

    if got != want {
        t.Errorf("got %g want %g,", want, got)
    }
}

//Testing Sqaure Perimeter
func TestSquare_Perimeter(t *testing.T) {
    _, result := Square(7.0)
    got := result
    want := 28.0

    if got != want {
        t.Errorf("got %g want %g,", want, got)
    }
}