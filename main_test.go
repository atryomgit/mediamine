package main

import (
	"reflect"
	"testing"
)

func TestCalculateCombinations(t *testing.T) {
	tests := []struct {
		amount    int
		banknotes []int
		expected  [][]int
	}{
		{
			amount:    400,
			banknotes: []int{5000, 2000, 1000, 500, 200, 100, 50},
			expected: [][]int{
				{200, 200},
				{200, 100, 100},
				{200, 100, 50, 50},
				{200, 50, 50, 50, 50},
				{100, 100, 100, 100},
				{100, 100, 100, 50, 50},
				{100, 100, 50, 50, 50, 50},
				{100, 50, 50, 50, 50, 50, 50},
				{50, 50, 50, 50, 50, 50, 50, 50},
			},
		},
		{
			amount:    100,
			banknotes: []int{50, 20, 10},
			expected: [][]int{
				{50, 50},
				{50, 20, 20, 10},
				{50, 20, 10, 10, 10},
				{50, 10, 10, 10, 10, 10},
				{20, 20, 20, 20, 20},
				{20, 20, 20, 20, 10, 10},
				{20, 20, 20, 10, 10, 10, 10},
				{20, 20, 10, 10, 10, 10, 10, 10},
				{20, 10, 10, 10, 10, 10, 10, 10, 10},
				{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			},
		},
		{
			amount:    0,
			banknotes: []int{50, 20, 10},
			expected:  [][]int{{}},
		},
	}

	for _, test := range tests {
		result := calculateCombinations(test.amount, test.banknotes)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("amount: %d, banknotes: %v, expected %v, got %v", test.amount, test.banknotes, test.expected, result)
		}
	}
}
