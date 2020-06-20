package main

import (
	"algorithms/stack"
	"fmt"
)

func main() {
	arrStack := stack.New("array")
	arrStack.Push(1)
	arrStack.Push(2)
	arrStack.Push(3)
	val, err := arrStack.Pop()

	listStack := stack.New("list")
	listStack.Push(1)
	listStack.Push(2)
	listStack.Push(3)
	val2, err2 := listStack.Pop()

	fmt.Println(val, err)
	fmt.Println(val2, err2)
}
