package main

func main() {
	//arr := createQuickUnion(5)
	//arr.connect(0, 2)
	//arr.connect(1, 4)
	//arr.connect(2, 3)
	//
	//arr2 := createWeighedQuickUnion(5)
	//arr2.connect(0, 2)
	//arr2.connect(1, 4)
	//arr2.connect(2, 3)
	//
	//fmt.Println(arr.isConnected(1, 4), arr2.isConnected(1, 4))
	//fmt.Println(arr.isConnected(0, 3), arr2.isConnected(0, 3))
	//fmt.Println(arr.isConnected(0, 1), arr2.isConnected(0, 1))

	arr3 := createWeighedQuickUnion(10)
	arr3.connect(3, 4)
	arr3.connect(4, 8)
	arr3.connect(6, 5)
	arr3.connect(9, 4)
	arr3.connect(2, 1)
	arr3.connect(5, 0)
	arr3.connect(7, 2)
	arr3.connect(6, 1)
	arr3.connect(7, 3)
	arr3.isConnected(7, 3)
}
