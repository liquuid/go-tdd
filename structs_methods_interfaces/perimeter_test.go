package smi

import "testing"

func TestPerimeter(t *testing.T){
	t.Run("Calculate perimeter 10x10", func(t *testing.T){
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got , want )
		}
	})

	t.Run("Calculate perimeter 15x10", func(t *testing.T){
		rectangle := Rectangle{15.0, 10.0}
		got := Perimeter(rectangle)
		want := 50.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got , want )
		}
	})

}

func TestArea(t *testing.T){
	areaTests := []struct {
		name string
		shape Shape
		want float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.50},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests{
		got := tt.shape.Area()
		if got != tt.want{
			t.Errorf("%#v got %.2f want %.2f",tt.shape, got , tt.want )
		}
	}
}