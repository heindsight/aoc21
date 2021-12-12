package stack

import "errors"

type Stack struct {
	stack []interface {}
}

var EmptyStackError = errors.New("The stack is empty!")

func MakeStack(capacity int) *Stack {
	return &Stack{stack: make([]interface {}, 0, capacity)}
}

func (s *Stack) Push(val interface {}) {
	s.stack = append(s.stack, val)
}

func (s *Stack) Peek() (interface {}, error) {
	if len(s.stack) == 0 {
		return nil, EmptyStackError
	}
	return s.stack[len(s.stack)-1], nil
}

func (s *Stack) Pop() (interface {}, error) {
	if len(s.stack) == 0 {
		return nil, EmptyStackError
	}
	value := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack) - 1]
	return value, nil
}

func (s *Stack) Length() int {
	return len(s.stack)
}
