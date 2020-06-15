package main

import (
	"algorithms/array"
	"fmt"
)

func main() {
	x := array.New(2)
	fmt.Println(x)
	v, err := x.Pop()
	fmt.Println(v, err)

	x.Push(1)
	x.Push(2)
	fmt.Println(x)
	x.Push(3)
	fmt.Println(x)
	x.Push(4)
	fmt.Println(x)
	x.Prepend(0)
	fmt.Println(x)
	v, err = x.Pop()
	fmt.Println(v, err, x)
	err = x.Delete(100)
	fmt.Println(err)
	err = x.Delete(1)
	fmt.Println(err)
	x.Push(2)
	x.Push(3)
	x.Push(2)
	x.Remove(2)
	x.Remove(0)
	fmt.Println(x)
	i, err := x.Find(3)
	fmt.Println(i, err)
	i, err = x.Find(0)
	fmt.Println(i, err)

	err = x.Insert(0, 99)
	fmt.Println(x, err)
	err = x.Insert(3, 30)
	fmt.Println(x, err)
	err = x.Insert(4, 40)
	fmt.Println(x, err)
	err = x.Insert(99, 0)
	fmt.Println(x, err)

}
