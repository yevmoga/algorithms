package main

func main() {
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
