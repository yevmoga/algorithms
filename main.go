package main

import (
	"fmt"
)

func main() {
	arr := createQuickUnion(5)
	arr.connect(0, 2)
	arr.connect(1, 4)
	arr.connect(2, 3)

	fmt.Println(arr.isConnected(1, 4))
	fmt.Println(arr.isConnected(0, 3))
	fmt.Println(arr.isConnected(0, 1))
}
