package stack

import "errors"

type Stack interface {
	Push(interface{})
	Peek() (interface{}, error)
	Pop() (interface{}, error)
	Length() int
}

type stack struct {
	stack []interface{}
}

var EmptyStackError = errors.New("The stack is empty!")

func NewStack(capacity int) *stack {
	return &stack{stack: make([]interface{}, 0, capacity)}
}

func (s *stack) Push(val interface{}) {
	s.stack = append(s.stack, val)
}

func (s *stack) Peek() (interface{}, error) {
	if len(s.stack) == 0 {
		return nil, EmptyStackError
	}
	return s.stack[len(s.stack)-1], nil
}

func (s *stack) Pop() (interface{}, error) {
	if len(s.stack) == 0 {
		return nil, EmptyStackError
	}
	value := s.stack[len(s.stack)-1]
	s.stack[len(s.stack)-1] = nil
	s.stack = s.stack[:len(s.stack)-1]
	return value, nil
}

func (s *stack) Length() int {
	return len(s.stack)
}
