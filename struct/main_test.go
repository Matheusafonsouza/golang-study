package main

import "testing"

func TestArea(t *testing.T) {
	verifyMessage := func(t *testing.T, result float64, expected float64) {
		t.Helper()
		if expected != result {
			t.Errorf("Result %.2f, expected %.2f", result, expected)
		}
	}

	areaTests := []struct {
		form     Form
		expected float64
	}{
		{Retangle{height: 10.0, width: 10.0}, 100.0},
		{Circle{radius: 10.0}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		result := tt.form.area()
		verifyMessage(t, result, tt.expected)
	}
}

func TestPerimeter(t *testing.T) {
	verifyMessage := func(t *testing.T, result float64, expected float64) {
		t.Helper()
		if expected != result {
			t.Errorf("Result %.2f, expected %.2f", result, expected)
		}
	}

	areaTests := []struct {
		form     Form
		expected float64
	}{
		{Retangle{height: 10.0, width: 10.0}, 40.0},
	}

	for _, tt := range areaTests {
		result := tt.form.perimeter()
		verifyMessage(t, result, tt.expected)
	}
}
