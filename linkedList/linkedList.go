package linkedList

import (
	"errors"
	"fmt"
)

type single struct {
	head  *single
	value int
	size  uint
	//tail  *single
}

func New() *single {
	return &single{}
}

func (s *single) Size() uint {
	return s.size
}

func (s *single) Empty() bool {
	return s.size == 0
}

func (s *single) Print(withValue bool) {
	if withValue {
		fmt.Print(s.value, ":", s.size, " | ")
	} else {
		fmt.Print(s.value, " | ")
	}
	if s.head != nil {
		s.head.Print(withValue)
	} else {
		fmt.Println()
	}
}

func (s *single) ValueAt(at int) (int, error) {
	if at < 1 || at > int(s.size) {
		return -1, errors.New("incorrect position")
	}

	x := s
	i := 0
	for i < at {
		x = x.head
		i++
	}

	if x == nil {
		return -1, errors.New("position not found")
	}

	return x.value, nil
}

func (s *single) PushFront(value int) {
	s.size++
	newLinkedList := New()
	newLinkedList.value = value
	newLinkedList.head = s.head
	newLinkedList.size = s.size - 1
	s.head = newLinkedList
}

func (s *single) PopFront() (int, error) {
	if s.size < 1 {
		return -1, errors.New("no elements")
	}
	value := s.head.value
	s.size--
	s.head = s.head.head

	return value, nil
}

func (s *single) PushBack(value int) {
	s.size++
	x := s
	for x.head != nil {
		x = x.head
		x.size++
	}

	x.head = New()
	x.head.value = value
}

func (s *single) PopBack() (int, error) {
	if s.size < 1 {
		return -1, errors.New("no elements")
	}

	x := s
	for x.head.head != nil {
		x.size--
		x = x.head
	}

	value := x.head.value
	x.size--
	x.head = nil

	return value, nil
}

func (s *single) Front() (int, error) {
	return s.ValueAt(1)
}

func (s *single) Back() (int, error) {
	return s.ValueAt(int(s.size))
}

func (s *single) Insert(index, value int) error {
	if s.size < 1 {
		return errors.New("no elements")
	}
	if index < 0 || uint(index) > s.size {
		return errors.New("incorrect index")
	}

	if index == 0 {
		s.PushFront(value)
		return nil
	}
	if index == int(s.size) {
		s.PushBack(value)
		return nil
	}

	s.size++

	tmp, i := s, 0
	for i < index {
		tmp = tmp.head
		tmp.size++
		i++
	}

	current := New()
	current.value = value
	current.size = tmp.size - 1
	current.head = tmp.head
	tmp.head = current

	return nil
}

func (s *single) Erase(index int) error {
	if s.size < 1 {
		return errors.New("no elements")
	}
	if index < 0 || uint(index) > s.size {
		return errors.New("incorrect index")
	}

	if index == 0 {
		_, err := s.PopFront()
		return err
	}
	if index == int(s.size) {
		_, err := s.PopBack()
		return err
	}

	s.size--

	tmp, i := s, 0
	for i < index {
		tmp = tmp.head
		tmp.size--
		i++
	}

	tmp.head = tmp.head.head

	return nil
}

//(n) - returns the value of the node at nth position from the end of the list
func (s *single) ValueFromEnd(n int) (int, error) {
	if s.size < 1 {
		return -1, errors.New("no elements")
	}
	if n < 0 || uint(n) > s.size {
		return -1, errors.New("incorrect index")
	}

	index := int(s.size) - n
	tmp, i := s, 0
	for i != index {
		tmp = tmp.head
		i++
	}

	return tmp.value, nil
}

// removes the first item in the list with this value
func (s *single) RemoveByValue(value int) error {
	if s.size < 1 {
		return errors.New("no elements")
	}

	tmp := s
	for tmp.head != nil && tmp.head.value != value {
		tmp = tmp.head
	}

	if tmp.head == nil || tmp.head.value != value {
		return errors.New("can't find this element :(")
	}

	tmp.head = tmp.head.head

	tmp2 := s
	for tmp2.head != tmp.head {
		tmp2.size--
		tmp2 = tmp2.head
	}
	tmp2.size--

	return nil
}

func (s *single) Reverse() {
	var (
		i       uint = 0
		prev    *single
		current = s.head
		next    *single
	)

	for current != nil {
		next = current.head
		current.head = prev
		prev = current
		current = next
		prev.size = i
		i++
	}
	s.head = prev
}
