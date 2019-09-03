package bubblesort

func Sort(nums []int) []int {
	size := len(nums)
	swaped := true

	for swaped {
		swaped = false

		for i := 0; i < size-1; i++ {
			current := nums[i]
			next := nums[i+1]

			if next < current {
				nums[i] = next
				nums[i+1] = current

				swaped = true
			}
		}

	}

	return nums
}
