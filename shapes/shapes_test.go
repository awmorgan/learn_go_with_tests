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
		shape Shape
		want float64
	} {
		{Rectangle{12,6},72},
		{Circle{10}, 314.1592653589793},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// }

	// t.Run("rectangles", func(t *testing.T) {
	// 	r := Rectangle{12, 6}
	// 	checkArea(t, r, 72)
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	c := Circle{10}
	// 	checkArea(t, c, 314.1592653589793)
	// })
}
