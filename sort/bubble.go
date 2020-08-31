package sort

func Bubble(arr []int) {

	l := len(arr)
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {

			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}

			if j == l {
				l--
			}
		}
	}

}
