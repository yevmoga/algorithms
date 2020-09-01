package sort

func Merge(arr []int) {
	newArr := merge(arr)
	for i, el := range newArr {
		arr[i] = el
	}
}

func merge(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	middle := len(arr) / 2

	left := merge(arr[:middle])
	right := merge(arr[middle:])

	index, lIndex, rIndex, val := 0, 0, 0, 0
	maxLen := len(left) + len(right)
	res := make([]int, maxLen)

	for index < maxLen {
		if lIndex < len(left) && rIndex >= len(right) || lIndex < len(left) && left[lIndex] <= right[rIndex] {
			val = left[lIndex]
			lIndex++
		} else {
			val = right[rIndex]
			rIndex++
		}

		res[index] = val
		index++
	}

	return res
}
