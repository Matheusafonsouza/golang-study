package main

import "testing"

func TestSum(t *testing.T) {
	verifyMessage := func(t *testing.T, result int, expected int) {
		t.Helper()
		if expected != result {
			t.Errorf("Result %d, expected %d", result, expected)
		}
	}

	t.Run("Should sum all list numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		result := sum(numbers)
		expected := 15
		verifyMessage(t, result, expected)
	})
}
