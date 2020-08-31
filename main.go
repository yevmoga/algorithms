package main

import (
	"algorithms/sort"
	"fmt"
)

func main() {
	arr := []int{100, 3, -19, 99, 1, 17, 99}

	sort.Bubble(arr)

	fmt.Println(arr)
}
