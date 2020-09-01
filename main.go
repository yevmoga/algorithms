package main

import (
	"algorithms/sort"
	"fmt"
)

func main() {
	//arr := []int{5, 4, 3, 2, 1}
	arr := []int{100, 3, -19, 99, 1, 17, 99}

	sort.Insertion(arr)

	fmt.Println(arr)
}
