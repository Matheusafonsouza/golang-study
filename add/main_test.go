package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	verifyMessage := func(t *testing.T, result int, expected int) {
		t.Helper()
		if expected != result {
			t.Errorf("Result %d, expected %d", result, expected)
		}
	}

	t.Run("Should add two numbers", func(t *testing.T) {
		sum := add(2, 2)
		expected := 4
		verifyMessage(t, sum, expected)
	})
}
