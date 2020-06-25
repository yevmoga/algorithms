package hash

import (
	"errors"
)

type item struct {
	next  *item
	key   int
	value int
}

type table struct {
	data []*item
	size int
}

func New(size int) *table {
	if size <= 0 {
		panic("incorrect size")
	}

	return &table{
		data: make([]*item, size),
		size: size,
	}
}

func (t *table) Insert(key, value int) {
	k := t.hash(key)

	i := &item{
		key:   key,
		value: value,
	}

	if t.data[k] == nil {
		t.data[k] = i
	} else {
		tmp := t.data[k]
		for tmp.next != nil {
			tmp = tmp.next
		}
		tmp.next = i
	}
}

func (t *table) Delete(key int) {
	k := t.hash(key)

	i := t.data[k]
	if i == nil {
		return
	}

	if i.key == key {
		if i.next == nil {
			t.data[k] = nil
		} else {
			t.data[k] = i.next
		}
		return
	}

	for i != nil && i.next != nil && i.next.key != key {
		i = i.next
	}

	if i != nil && i.next.key == key {
		*i.next = *i.next.next
		return
	}
}

func (t *table) Search(key int) (int, error) {
	k := t.hash(key)

	i := t.data[k]
	for i != nil && i.key != key && i.next != nil {
		i = i.next
	}

	if i != nil && i.key == key {
		return i.value, nil
	}

	return 0, errors.New("element not found")
}

func (t *table) hash(key int) int {
	if key < 0 {
		key = -key
	}
	return key % t.size
}
