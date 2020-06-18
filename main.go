package main

import (
	"algorithms/linkedList"
	"fmt"
)

func main() {
	l := linkedList.New()

	fmt.Println(l.Size())
	fmt.Println(l.Empty())

	l.PushFront(4)
	l.PushFront(3)
	l.PushFront(2)
	l.PushFront(1)

	fmt.Println()

	fmt.Println(l.ValueAt(0))
	fmt.Println(l.ValueAt(1))
	fmt.Println(l.ValueAt(2))
	fmt.Println(l.ValueAt(3))
	fmt.Println(l.ValueAt(4))
	fmt.Println(l.ValueAt(5))
	fmt.Println(l.ValueAt(100))

	l.Print(false)
	fmt.Println(l.PopFront())
	l.Print(false)

	fmt.Println()

	l.PushFront(1)
	l.Print(true)
	l.PushBack(98)
	l.PushBack(99)
	l.PushBack(100)
	l.Print(true)
	popBack, err := l.PopBack()
	l.Print(true)
	fmt.Print(popBack, err)
	fmt.Println(l.PopBack())
	fmt.Println(l.PopBack())
	fmt.Println(l.PopBack())
	fmt.Println(l.PopBack())
	fmt.Println(l.PopBack())
	fmt.Println(l.PopBack())
	fmt.Println(l.PopBack())
	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)
	l.PushBack(40)
	l.PushBack(50)
	l.Print(true)
	fmt.Println(l.Front())
	fmt.Println(l.Back())

	fmt.Println(l.Insert(0, 5))
	fmt.Println(l.Insert(2, 15))
	fmt.Println(l.Insert(7, 100))
	fmt.Println(l.Insert(400, 400))
	l.Print(true)

	fmt.Println()
	fmt.Println(l.Erase(0))
	fmt.Println(l.Erase(3))
	fmt.Println(l.Erase(5))
	fmt.Println(l.Erase(7))
	l.Print(true)

	fmt.Println(l.ValueFromEnd(2))

	l.Print(true)
	l.RemoveByValue(20)
	fmt.Println(l.RemoveByValue(999))
	l.Print(true)

	fmt.Println()
	fmt.Println()
	l.Print(true)
	l.Reverse()
	l.Print(true)
	l.Reverse()
	l.Print(true)
}
