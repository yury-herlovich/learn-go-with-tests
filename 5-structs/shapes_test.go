package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	expected := 40.0

	if got != expected {
		t.Errorf("expected %f but got %f", expected, got)
	}
}

func TestArea(t *testing.T) {

	// Table driven tests
	areaTests := []struct {
		shape    Shape
		expected float64
	}{
		{shape: Rectangle{12.0, 6.0}, expected: 72.0},
		{shape: Circle{10.0}, expected: 314.1592653589793},
		{shape: Triangle{12, 6}, expected: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		if got != tt.expected {
			t.Errorf("%#v expected %g but got %g", tt.shape, tt.expected, got)
		}
	}

	// regular tests
	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		got := shape.Area()

		if got != expected {
			t.Errorf("expected %g but got %g", expected, got)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}
