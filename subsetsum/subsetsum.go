package subsetsum

import (
	"sort"
)

func Exists(nums []uint, sum uint) bool {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for i, numA := range nums {
		if numA >= sum {
			continue
		}

		if numA < sum/2 {
			break
		}

		for _, numB := range nums[i+1:] {
			currentSum := numA + numB
			if sum == currentSum {
				return true
			}

			if currentSum < currentSum {
				break
			}

		}
	}

	return false
}

func DExists(nums []uint, sum uint) bool {
	numsCache := map[uint]bool{}

	for _, num := range nums {
		if numsCache[num] {
			return true
		}

		rest := sum - num
		numsCache[rest] = true
	}

	return false
}
