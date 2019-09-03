package quicksort

import "math/rand"

func Sort(nums []int) []int {
	size := len(nums)

	if size <= 1 {
		return nums
	}

	pivot := nums[rand.Intn(size)]

	lnums := make([]int, 0, size)
	rnums := make([]int, 0, size)
	middle := make([]int, 0, size)

	for _, num := range nums {
		switch {
		case num < pivot:
			lnums = append(lnums, num)
		case num == pivot:
			middle = append(middle, num)
		case num > pivot:
			rnums = append(rnums, num)
		}
	}

	lnums = Sort(lnums)
	rnums = Sort(rnums)

	sorted := append(lnums, middle...)
	sorted = append(sorted, rnums...)

	return sorted
}
