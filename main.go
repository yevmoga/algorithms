package main

import (
	"algorithms/hash"
	"fmt"
)

func main() {
	h := hash.New(10)

	h.Insert(100, 1)
	h.Insert(101, 2)
	h.Insert(102, 3)

	h.Insert(110, 4)
	h.Insert(111, 5)
	h.Insert(112, 6)

	h.Insert(121, 7)

	fmt.Println(h.Search(100))
	fmt.Println(h.Search(101))
	fmt.Println(h.Search(110))
	fmt.Println(h.Search(999))

	h.Delete(100)
	h.Delete(110)
	h.Delete(111)
	h.Delete(999)
}
