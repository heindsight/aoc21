package solutions

import (
	"errors"
	"fmt"
	"io"

	"github.com/heindsight/aoc21/registry"
)

type window struct {
	values []int
	head   int
	tail   int
	count  int
}

func newWindow(capacity int) *window {
	return &window{
		values: make([]int, capacity, capacity),
		head:   0,
		tail:   0,
		count:  0,
	}
}

func (win *window) pop() (int, error) {
	if win.count == 0 {
		return -1, errors.New("Cannot pop from empty window")
	}

	value := win.values[win.head]

	win.count--
	win.head = (win.head + 1) % cap(win.values)

	return value, nil
}

func (win *window) push(value int) error {
	if win.is_full() {
		return errors.New("Cannot push, window is full")
	}

	if win.count > 0 {
		win.tail = (win.tail + 1) % cap(win.values)
	}
	win.values[win.tail] = value
	win.count++

	return nil
}

func (win *window) is_full() bool {
	return win.count == cap(win.values)
}

type day01bSolution struct {
	depth_increases int
	depth_window    *window
}

func makeDay01bSolution() day01bSolution {
	return day01bSolution{
		depth_increases: 0,
		depth_window: newWindow(3),
	}
}

func (soln day01bSolution) Solve() error {
	for {
		var depth int
		_, err := fmt.Scan(&depth)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if soln.depth_window.is_full() {
			prev_depth, _ := soln.depth_window.pop()

			if depth > prev_depth {
				soln.depth_increases++
			}

		}
		soln.depth_window.push(depth)
	}

	fmt.Println(soln.depth_increases)
	return nil
}

func init() {
	soln := makeDay01bSolution()
	if err := registry.RegisterSolution("day01b", soln); err != nil {
		fmt.Println("Failed to register day01b solution", err)
	}
}
