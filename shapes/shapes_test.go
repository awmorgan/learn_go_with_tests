package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	type args struct {
		r Rectangle
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"10x10", args{r: Rectangle{10, 10}}, 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Perimeter(tt.args.r); got != tt.want {
				t.Errorf("Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{width: 12, height: 6}, want: 72},
		{name: "circle", shape: Circle{radius: 10}, want: 314.1592653589793},
		{name: "triangle", shape: Triangle{base: 12, height: 6}, want: 36},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		})
	}
}
