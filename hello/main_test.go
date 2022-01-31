package main

import "testing"

func TestHello(t *testing.T) {
	verifyMessage := func(t *testing.T, result string, expected string) {
		t.Helper()
		if expected != result {
			t.Errorf("Result %s, expected %s", result, expected)
		}
	}

	t.Run("Show say hello world if doesnt have someone", func(t *testing.T) {
		result := hello("", "")
		expected := "Hello, world"
		verifyMessage(t, result, expected)
	})

	t.Run("Should say hello to someone", func(t *testing.T) {
		result := hello("Afonso", "")
		expected := "Hello, Afonso"
		verifyMessage(t, result, expected)
	})

	t.Run("Should say hello in spanish", func(t *testing.T) {
		result := hello("Afonso", "spanish")
		expected := "Hola, Afonso"
		verifyMessage(t, result, expected)
	})

	t.Run("Should say hello in english", func(t *testing.T) {
		result := hello("Afonso", "english")
		expected := "Hello, Afonso"
		verifyMessage(t, result, expected)
	})

	t.Run("Should say hello in french", func(t *testing.T) {
		result := hello("Afonso", "french")
		expected := "Bonjour, Afonso"
		verifyMessage(t, result, expected)
	})
}
