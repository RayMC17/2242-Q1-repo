package main

import (
	"testing"
)

// Testing Triangle Area
func TestArea(t *testing.T) {
	t.Run("triangle", func(t *testing.T) {
		triangle := triangle{10.5, 6.6}
		got := triangle.area()
		want := 34.65

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}

	})
}

//Testing Triangle Perimeter 
func TestPerimeter(t *testing.T) {
	t.Run("triangle", func(t *testing.T) {
		triangle := triangle{10.5, 6.6}
		got := triangle.perimeter()
		want := 12.402015965156632

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
