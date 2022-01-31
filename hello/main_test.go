package main

import "testing"

func TestHello(t *testing.T) {
	result := hello("Afonso")
	expected := "Hello, Afonso!"

	if expected != result {
		t.Errorf("Result %s, expected %s", result, expected)
	}
}
