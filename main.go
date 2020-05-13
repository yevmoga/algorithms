package main

import (
	"fmt"
)

func main() {
	arr := create(5)
	arr.connect(0, 2)
	arr.connect(1, 4)
	arr.connect(2, 3)

	fmt.Println(arr.isConnected(1,4))
	fmt.Println(arr.isConnected(0,3))
	fmt.Println(arr.isConnected(0,1))
}

type QuickFindUF []int

func create(n int) QuickFindUF {
	arr := make(QuickFindUF, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

func (arr QuickFindUF) connect(a, b int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == a {
			arr[i] = b
		}
	}
}

func (arr QuickFindUF) isConnected(a, b int) bool {
	return arr[a] == arr[b]
}
