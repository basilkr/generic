package stack

import (
	"testing"
)

func TestStack0(t *testing.T) {
	st0 := Make[int](2)
	if _, err := st0.Pop(); err == nil {
		t.Errorf("Should get error popping empty stack")
	}
	if _, err := st0.Peek(); err == nil {
		t.Errorf("Should get error peeking empty stack")
	}
	st0.Push(1).Push(2).Push(3).Push(4).Push(5).Push(6)
	if st0.Depth() != 6 {
		t.Errorf("Expected depth 6 ")
	}
	if x, err := st0.Peek(); err != nil {
		t.Errorf("Should not get error peeking non-empty stack")
	} else if x != 6 {
		t.Errorf("Expected the top value to be 6")
	}
	if x, err := st0.Pop(); err != nil {
		t.Errorf("Should not get error popping non-empty stack")
	} else if x != 6 {
		t.Errorf("Expected the top value to be 6")
	}
	if x, err := st0.Pop(); err != nil {
		t.Errorf("Should not get error popping non-empty stack")
	} else if x != 5 {
		t.Errorf("Expected the top value to be 5")
	}
}
