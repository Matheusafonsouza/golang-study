package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	verifyMessage := func(t *testing.T, result int, expected int) {
		t.Helper()
		if expected != result {
			t.Errorf("Result %d, expected %d", result, expected)
		}
	}

	t.Run("Should sum all slice numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := sum(numbers)
		expected := 6
		verifyMessage(t, result, expected)
	})
}

func TestSumAll(t *testing.T) {
	verifyMessage := func(t *testing.T, result []int, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Result %d, expected %d", result, expected)
		}
	}

	t.Run("Should sum all slices in one slice", func(t *testing.T) {
		result := sumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}
		verifyMessage(t, result, expected)
	})
}
