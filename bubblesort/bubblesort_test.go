package bubblesort_test

import (
	"math/rand"
	"testing"

	"github.com/alesshh/algorithms/bubblesort"
	"github.com/google/go-cmp/cmp"
)

func TestSort(t *testing.T) {
	scenarios := []struct {
		given    []int
		expected []int
	}{
		{
			given:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			given:    []int{5},
			expected: []int{5},
		},
		{
			given:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			given:    []int{5, 1, 2, 8, 3},
			expected: []int{1, 2, 3, 5, 8},
		},
		{
			given:    []int{4, 1, 2, 8, 3, 3},
			expected: []int{1, 2, 3, 3, 4, 8},
		},
	}

	for _, scenario := range scenarios {
		sorted := bubblesort.Sort(scenario.given)

		if !cmp.Equal(scenario.expected, sorted) {
			t.Error(cmp.Diff(scenario.expected, sorted))
		}
	}
}

func BenchmarkSort(b *testing.B) {
	b.StopTimer()

	nums := make([]int, b.N)

	for i := 0; i < b.N; i++ {
		nums[i] = rand.Intn(b.N)
	}

	b.StartTimer()

	bubblesort.Sort(nums)
}
