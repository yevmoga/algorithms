package stack

import (
	"algorithms/array"
	"algorithms/linkedList"
)

type stackable interface {
	PushBack(int)
	PopBack() (int, error)
}

type stack struct {
	data stackable
}

func New(t string) *stack {
	var data stackable
	switch t {
	case "array":
		data = array.New(1)
	default:
		data = linkedList.New()
	}

	return &stack{
		data: data,
	}
}

func (s *stack) Push(value int) {
	s.data.PushBack(value)
}

func (s *stack) Pop() (int, error) {
	return s.data.PopBack()
}
