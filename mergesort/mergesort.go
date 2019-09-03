package mergesort

func Sort(nums []int) []int {
	size := len(nums)
	if size <= 1 {
		return nums
	}

	middle := size / 2
	lnums := nums[:middle]
	rnums := nums[middle:]

	return merge(Sort(lnums), Sort(rnums))
}

func merge(lnums, rnums []int) []int {
	lsize := len(lnums)
	rsize := len(rnums)
	sorted := make([]int, 0, lsize+rsize)

	var lindex, rindex int

	for lindex < lsize && rindex < rsize {
		if lnums[lindex] < rnums[rindex] {
			sorted = append(sorted, lnums[lindex])
			lindex++
		} else {
			sorted = append(sorted, rnums[rindex])
			rindex++
		}
	}

	sorted = append(sorted, lnums[lindex:]...)
	sorted = append(sorted, rnums[rindex:]...)

	return sorted
}
