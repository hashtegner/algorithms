package subsetsum_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/alesshh/algorithms/subsetsum"
)

func TestExists(t *testing.T) {
	tests := []struct {
		nums     []uint
		sum      uint
		expected bool
	}{

		{[]uint{}, 10, false},
		{[]uint{5}, 6, false},
		{[]uint{9}, 9, false},
		{[]uint{9, 1}, 9, false},
		{[]uint{5, 1, 1, 1, 1}, 9, false},
		{[]uint{1, 5}, 6, true},
		{[]uint{1, 3, 5, 7, 6, 6, 2, 4}, 9, true},
		{[]uint{8, 7, 6, 5, 4, 3, 2, 1}, 9, true},
		{[]uint{8, 1, 6, 5, 4, 3, 2}, 9, true},
		{[]uint{4, 1, 6, 5, 4, 3}, 8, true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Exists(nums: %v, sum: %v)", test.nums, test.sum), func(t *testing.T) {
			result := subsetsum.Exists(test.nums, test.sum)
			if result != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, result)
			}
		})

		t.Run(fmt.Sprintf("DExists(nums: %v, sum: %v)", test.nums, test.sum), func(t *testing.T) {
			result := subsetsum.DExists(test.nums, test.sum)
			if result != test.expected {
				t.Fatalf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func BenchmarkExists(b *testing.B) {
	b.StopTimer()

	sum := uint(rand.Intn(b.N))
	nums := make([]uint, b.N)

	for i := 0; i < b.N; i++ {
		nums[i] = uint(rand.Intn(b.N))
	}

	b.StartTimer()

	subsetsum.Exists(nums, sum)
}

func BenchmarkDExists(b *testing.B) {
	b.StopTimer()

	sum := uint(rand.Intn(b.N))
	nums := make([]uint, b.N)

	for i := 0; i < b.N; i++ {
		nums[i] = uint(rand.Intn(b.N))
	}

	b.StartTimer()

	subsetsum.DExists(nums, sum)
}
