package stack

import (
	"errors"
)

type stack[T any] struct {
	impl []T
}

// Make creates an empty Stack[T] with a given initial capacity
func Make[T any](capacity int) stack[T] {
	return stack[T]{impl: make([]T, 0, capacity)}
}

// Push a value onto the top of the Stack
func (it *stack[T]) Push(val T) *stack[T] {
	it.impl = append(it.impl, val)
	return it
}

// View the top item on the Stack
func (it stack[T]) Peek() (T, error) {
	if len(it.impl) == 0 {
		var un_set T
		return un_set, errors.New("peeking empty Stack")
	}
	return it.impl[len(it.impl)-1], nil
}

// Pop the top item of the Stack and return it
func (it *stack[T]) Pop() (T, error) {
	if len(it.impl) == 0 {
		var un_set T
		return un_set, errors.New("popping empty hStack")
	}
	value := it.impl[len(it.impl)-1]
	it.impl = it.impl[:len(it.impl)-1]
	it.impl = reclaim(it.impl)
	return value, nil
}

// PopIt consumes the error sets res to nil on error,
// sets res to the popped value, returns pointer to same stack
func (it *stack[T]) PopIt(res *T) *stack[T] {
	x, err := it.Pop()
	if err != nil {
		res = nil
	} else {
		*res = x
	}
	return it
}

// Return the number of items in the Stack
func (it stack[T]) Depth() int {
	return len(it.impl)
}

func reclaim[T any](slice []T) []T {
	// drop unused memory if capacity is excessive relative to length
	// keep capacity between 2 and 3 times the length
	if cap(slice) > 3*len(slice) {
		tmp := make([]T, len(slice), 2*len(slice))
		copy(tmp, slice)
		return tmp
	}
	return slice
}
