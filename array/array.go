package array

import (
	"errors"
)

type Array struct {
	val  []int
	size int
	cap  int
}

func New(cap int) *Array {
	if cap <= 0 {
		panic("cap <= 0")
	}

	arr := make([]int, cap, cap)
	return &Array{
		val:  arr,
		size: 0,
		cap:  cap,
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

	a.val[a.size] = value
	a.size++
}

func (a *Array) Insert(index, value int) error {
	if index > a.cap-1 {
		return errors.New("index is higher then size")
	}

	a.val[index] = value
	if index > a.size {
		a.size++
	}

	return nil
}

func (a *Array) Prepend(value int) {
	if a.size == a.cap {
		a.resize()
	}

	for i := a.size; i > 0; i-- {
		a.val[i], a.val[i-1] = a.val[i-1], a.val[i]
	}

	a.val[0] = value
	a.size++
}

func (a *Array) Pop() (int, error) {
	if a.size == 0 {
		return 0, errors.New("size = 0")
	}

	res := a.val[a.size-1]
	a.val[a.size-1] = 0
	a.size--

	return res, nil
}

func (a *Array) Delete(index int) error {
	if index > a.size {
		return errors.New("incorrect index")
	}

	for i := index; i < a.size; i++ {
		a.val[i] = a.val[i+1]
	}
	a.size--

	return nil
}

func (a *Array) Find(value int) (int, error) {
	return a.findAfterIndex(0, value)
}

func (a *Array) Remove(value int) {
	for i, err := a.findAfterIndex(0, value); err == nil; i, err = a.findAfterIndex(i, value) {
		a.Delete(i)
	}
}

func (a *Array) findAfterIndex(index, value int) (int, error) {
	for i := index; i < a.size; i++ {
		if a.val[i] == value {
			return i, nil
		}
	}
	return -1, errors.New("element not found")
}

func (a *Array) resize() {
	oldCap := a.cap
	newCap := oldCap * oldCap

	if newCap >= 1024 {
		newCap = oldCap + oldCap/4
	}

	a.cap = newCap
	arr := make([]int, a.cap, a.cap)
	for i := 0; i < a.size; i++ {
		arr[i] = a.val[i]
	}

	a.val = arr
}
