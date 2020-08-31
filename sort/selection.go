package sort

func Selection(arr []int) {

	leftPointer, length := 0, len(arr)

	for leftPointer != length {
		minPointer := leftPointer

		for i := leftPointer; i < length; i++ {
			if arr[i] < arr[minPointer] {
				minPointer = i
			}
		}

		arr[minPointer], arr[leftPointer] = arr[leftPointer], arr[minPointer]
		leftPointer++

	}

}
