package queue

import (
	"algorithms/array"
	"algorithms/linkedList"
)

type queueble interface {
	PushFront(int)
	PopBack() (int, error)
}

type queue struct {
	data queueble
}

func New(t string) *queue {
	var data queueble
	switch t {
	case "array":
		data = array.New(1)
	default:
		data = linkedList.New()
	}

	return &queue{
		data: data,
	}
}

func (q *queue) Push(value int) {
	q.data.PushFront(value)
}

func (q *queue) Pop() (int, error) {
	return q.data.PopBack()
}
