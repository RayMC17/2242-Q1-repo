package main

import (
	"testing"
)

//Testing area of circle
func TestArea(t *testing.T){
	t.Run("circle", func(t *testing.T){
		circle := circle{9}
		got := circle.area()
		want := 254.46900494077323

		if got !=want{
			t.Errorf("got %g want %g", got, want)
		}
	})
}

func TestPerimeter(t *testing.T){
	t.Run("circle", func(t *testing.T){
		circle := circle{9}
		got := circle.perimeter()
		want := 56.57142857142857

		if got !=want{
			t.Errorf("got %g want %g", got, want)
		}
	})
}

