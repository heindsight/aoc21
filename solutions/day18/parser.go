package day18

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/heindsight/aoc21/utils/stack"
)

var tokeniser = regexp.MustCompile(`[][,]|\d+`)

func Parse(str string) *snailNumber {
	state_machine := newSnailParseSM()

	for _, token := range tokeniser.FindAllString(str, -1) {
		state_machine.handleToken(token)
	}
	return state_machine.parsed
}

type snailState func(*snailParseSM, string)

type snailParseSM struct {
	stack stack.Stack
	state snailState
	parsed *snailNumber
}

func newSnailParseSM() *snailParseSM {
	sm := &snailParseSM{}
	sm.stack = stack.NewStack(128)
	sm.state = startState
	return sm
}

func (sm *snailParseSM) handleToken(token string) {
	sm.state(sm, token)
}

func startState(sm *snailParseSM, token string) {
	if token != "[" {
		unexpected(token)
	}

	sm.stack.Push(token)
	sm.state = pairOpened
}

func pairOpened(sm *snailParseSM, token string) {
	switch token {
	case "]", ",":
		unexpected(token)
	case "[":
		sm.stack.Push(token)
	default:
		sm.stack.Push(parseInt(token))
		sm.state = gotFirstNumber
	}
}

func gotFirstNumber(sm *snailParseSM, token string) {
	if token != "," {
		unexpected(token)
	}
	sm.state = awaitSecondNumber
}

func awaitSecondNumber(sm *snailParseSM, token string) {
	switch token {
	case "]", ",":
		unexpected(token)
	case "[":
		sm.stack.Push(token)
		sm.state = pairOpened
	default:
		sm.stack.Push(parseInt(token))
		sm.state = gotSecondNumber
	}
}

func gotSecondNumber(sm *snailParseSM, token string) {
	if token != "]" {
		unexpected(token)
	}

	number := snailNumber{elements: &[2]snailNumber{}}
	number.elements[1] = popSnailNumber(sm.stack)
	number.elements[0] = popSnailNumber(sm.stack)
	sm.stack.Pop()

	if sm.stack.Length() == 0 {
		sm.parsed = &number
		sm.state = done
		return
	}

	top, _ := sm.stack.Peek()
	switch top.(type) {
	case string:
		sm.stack.Push(number)
		sm.state = gotFirstNumber
	case snailNumber:
		sm.stack.Push(number)
		sm.state = gotSecondNumber
	default:
		panic(fmt.Errorf("Unexpected object of type %T on stack: %v\n", top, top))
	}
}

func done(sm *snailParseSM, token string) {
	unexpected(token)
}

func popSnailNumber(s stack.Stack) snailNumber {
	top, err := s.Pop()
	if err != nil {
		panic(err)
	}
	return top.(snailNumber)
}

func parseInt(token string) snailNumber {
	val, err := strconv.Atoi(token)
	if err != nil {
		panic(err)
	}
	return snailNumber{value: val}
}

func unexpected(token string) {
	panic("Unexpected token: " + token)
}
