package main

import "testing"

func TestRetangle(t *testing.T) {
	verifyMessage := func(t *testing.T, result float64, expected float64) {
		t.Helper()
		if expected != result {
			t.Errorf("Result %.2f, expected %.2f", result, expected)
		}
	}

	t.Run("Should calculate perimeter", func(t *testing.T) {
		retangle := Retangle{height: 10.0, width: 10.0}
		result := retangle.perimeter()
		expected := 40.0
		verifyMessage(t, result, expected)
	})

	t.Run("Should calculate area", func(t *testing.T) {
		retangle := Retangle{height: 10.0, width: 10.0}
		result := retangle.area()
		expected := 100.0
		verifyMessage(t, result, expected)
	})
}

func TestCircle(t *testing.T) {
	verifyMessage := func(t *testing.T, result float64, expected float64) {
		t.Helper()
		if expected != result {
			t.Errorf("Result %.2f, expected %.2f", result, expected)
		}
	}

	t.Run("Should calculate area", func(t *testing.T) {
		circle := Circle{radius: 10.0}
		result := circle.area()
		expected := 314.1592653589793
		verifyMessage(t, result, expected)
	})
}
