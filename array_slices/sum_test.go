package array_slices

import (
	"fmt"
	"reflect"
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


func TestSumAllTails(t *testing.T){
	checkSums := func(t *testing.T, got, want []int){
		t.Helper()
		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		result := SumAllTails([]int{1, 2, 3}, []int{2, 9, 10, 1}, []int{5, 5}, []int{4})
		expected := []int{5, 20, 5, 0}

		checkSums(t, result, expected)

	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		result := SumAllTails([]int{}, []int{2, 9, 10, 1})
		expected := []int{0, 20}

		checkSums(t, result, expected)
	})


}
