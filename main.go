package main

import (
	"algorithms/array"
	"fmt"
)

func main() {
	x := array.New(0, 2)
	fmt.Println(x)

	x.Push(1)
	x.Push(2)
	fmt.Println(x)
	x.Push(3)
	fmt.Println(x)
}
