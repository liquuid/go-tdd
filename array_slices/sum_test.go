package array_slices

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{1, 2, 3, 4}, 10},
	}

	for _, ts := range tests {
		t.Run(fmt.Sprintf("Array %v", ts.in), func(t *testing.T) {
			arr := ts.in
			result := Sum(arr)
			expected := ts.expected

			if result != expected {
				t.Errorf("Got %d, expected %d", result, expected)
			}
		})
	}
}
