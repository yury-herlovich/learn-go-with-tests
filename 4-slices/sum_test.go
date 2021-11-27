package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		sum := Sum(numbers)
		expected := 15

		if sum != expected {
			t.Errorf("expected %d but got %d", expected, sum)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("2 collections", func(t *testing.T) {
		sums := SumAll([]int{1, 2, 3}, []int{0, 7, 12, 76})
		expected := []int{6, 95}

		if !reflect.DeepEqual(sums, expected) {
			t.Errorf("expected %v but got %v", expected, sums)
		}
	})
}

func TestSumAllTrails(t *testing.T) {
	t.Run("slices with some numbers", func(t *testing.T) {
		sums := SumAllTrails([]int{2, 6, 17}, []int{3, 8, 14})
		expected := []int{23, 22}

		if !reflect.DeepEqual(sums, expected) {
			t.Errorf("expected %v but got %v", expected, sums)
		}
	})

	t.Run("empty slices", func(t *testing.T) {
		sums := SumAllTrails([]int{}, []int{3, 8, 14})
		expected := []int{0, 22}

		if !reflect.DeepEqual(sums, expected) {
			t.Errorf("expected %v but got %v", expected, sums)
		}
	})
}
