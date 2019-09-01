package shapes

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12.0, Height: 6.0}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12.0, Height: 6.0}, want: 36.0},
	}

	for _, tt := range areTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f %#v", got, tt.want, tt)
		}
	}
}

func ExamplePerimeter() {
	rectangle := Rectangle{10.0, 10.0}
	fmt.Println(Perimeter(rectangle))
	// Output: 40
}

func ExampleRectangle_Area() {
	rectangle := Rectangle{12.0, 6.0}
	fmt.Println(rectangle.Area())
	// Output: 72
}

func ExampleCircle_Area() {
	circle := Circle{10}
	fmt.Println(circle.Area())
	// Output: 314.1592653589793
}

func ExampleTriangle_Area() {
	triangle := Triangle{12.0, 6.0}
	fmt.Println(triangle.Area())
	// Output: 36
}
