package main

import "testing"

func TestLoop(t *testing.T) {
	verifyMessage := func(t *testing.T, result float64, expected float64) {
		t.Helper()
		if expected != result {
			t.Errorf("Result %.2f, expected %.2f", result, expected)
		}
	}

	t.Run("Should calculate perimeter", func(t *testing.T) {
		retangle := Retangle{height: 10.0, width: 10.0}
		result := perimeter(retangle)
		expected := 40.0
		verifyMessage(t, result, expected)
	})

	t.Run("Should calculate area", func(t *testing.T) {
		retangle := Retangle{height: 10.0, width: 10.0}
		result := area(retangle)
		expected := 100.0
		verifyMessage(t, result, expected)
	})
}
