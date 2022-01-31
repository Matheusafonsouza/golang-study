package main

import "testing"

func TestLoop(t *testing.T) {
	verifyMessage := func(t *testing.T, result string, expected string) {
		t.Helper()
		if expected != result {
			t.Errorf("Result %s, expected %s", result, expected)
		}
	}

	t.Run("Should repeat character", func(t *testing.T) {
		iterations := repeat("a")
		expected := "aaaaa"
		verifyMessage(t, iterations, expected)
	})
}
