package interfaces

import "testing"

// NOW THAT YOU HAVE SOME UNDERSTANDING OF STRUCTS WE CAN INTRODUCE "TABLE DRIVEN TESTS".

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		//* DIRECT VALUE OF STRUCT WITHOUT NAMING FIELDS
		// {Rectangle{12, 6}, 72.0},
		// {Circle{10}, 314.1592653589793},
		// {Triangle{12, 6}, 36.0},

		//*  OPTIONALLY NAME THE FIELDS
		{name: "rectangles", shape: Rectangle{width: 12, height: 6}, want: 72.0},
		{name: "circles", shape: Circle{radius: 10}, want: 314.1592653589793},
		{name: "triangles", shape: Triangle{base: 12, height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.area()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		})
	}
}
