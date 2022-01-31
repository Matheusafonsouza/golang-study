package main

import "testing"

func TestLoop(t *testing.T) {
	verifyMessage := func(t *testing.T, result string, expected string) {
		t.Helper()
		if expected != result {
			t.Errorf("Result %s, expected %s", result, expected)
		}
	}

	t.Run("Should repeat character N times", func(t *testing.T) {
		iterations := repeat("a", 6)
		expected := "aaaaaa"
		verifyMessage(t, iterations, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("a", 6)
	}
}
