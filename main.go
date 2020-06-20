package main

import (
	"algorithms/queue"
	"fmt"
)

func main() {
	arrQueue := queue.New("array")
	arrQueue.Push(1)
	arrQueue.Push(2)
	arrQueue.Push(3)
	val, err := arrQueue.Pop()

	listQueue := queue.New("list")
	listQueue.Push(1)
	listQueue.Push(2)
	listQueue.Push(3)
	val2, err2 := listQueue.Pop()

	fmt.Println(val, err)
	fmt.Println(val2, err2)
}
