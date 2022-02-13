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
		result := perimeter(10.0, 10.0)
		expected := 40.0
		verifyMessage(t, result, expected)
	})

	t.Run("Should calculate area", func(t *testing.T) {
		result := area(10.0, 10.0)
		expected := 100.0
		verifyMessage(t, result, expected)
	})
}
