package sort

func Insertion(arr []int) {
	for i := 1; i < len(arr); i++ {

		j, k := i, i-1
		for k >= 0 && arr[j] < arr[k] {
			arr[j], arr[k] = arr[k], arr[j]
			j--
			k--
		}
	}
}
