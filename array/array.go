package array

import (
	"errors"
)

type Array struct {
	val  []int
	size int
	cap  int
}

func New(size, cap int) *Array {
	if cap < size {
		cap = size
	}

	arr := make([]int, cap, cap)
	return &Array{
		val:  arr,
		cap:  cap,
		size: size,
	}
}

func (a *Array) Size() int {
	return a.size
}

func (a *Array) Cap() int {
	return a.cap
}

func (a *Array) IsEmpty() bool {
	return a.size == 0
}

func (a *Array) At(i int) (int, error) {
	if i < 0 || i > a.size {
		return 0, errors.New("index is out of range")
	}
	return a.val[i], nil
}

func (a *Array) Push(value int) {
	if a.size == a.cap {
		a.resize()
	}

	a.val[0] = value
	a.size++
}

func (a *Array) resize() {
	oldCap := a.cap
	newCap := oldCap * oldCap

	if newCap >= 1024 {
		newCap = oldCap + oldCap/4
	}

	a.cap = newCap
	arr := make([]int, a.size, a.cap)
	for i := 0; i < a.size; i++ {
		arr[i] = a.val[i]
	}

	a.val = arr
}
