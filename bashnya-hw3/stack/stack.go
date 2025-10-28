package stack

import (
	"errors"
)

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	top  *Node
	size int
}

func NewStack() *Stack {
	return &Stack{top: nil, size: 0}
}
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	value := s.top.value
	s.top = s.top.next
	s.size--
	return value, nil
}
func (s *Stack) Push(value int) {
	newNode := &Node{value: value, next: s.top}
	s.top = newNode
	s.size++
}
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}
func (s *Stack) Size() int {
	return s.size
}
func (s *Stack) Clear() {
	s.top = nil
	s.size = 0
}
