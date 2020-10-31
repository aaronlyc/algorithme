package data_struct

import "testing"

func TestSliceStack(t *testing.T) {
	stack := NewSliceStack(8)
	t.Logf("the stack item is: %v, the top point is: %d", stack.items, stack.i)

	stack.Pop()
	stack.Pop()
	stack.Pop()

	t.Logf("top is: %v", stack.Top())
	stack.Push(1)
	t.Logf("top is: %v", stack.Top())
	stack.Push(2)
	t.Logf("top is: %v", stack.Top())
	stack.Push(3)
	t.Logf("top is: %v", stack.Top())
	stack.Push(4)
	t.Logf("top is: %v", stack.Top())
	stack.Push(5)
	t.Logf("top is: %v", stack.Top())
	stack.Push(6)
	t.Logf("top is: %v", stack.Top())
	stack.Push(7)
	t.Logf("top is: %v", stack.Top())
	stack.Push(8)
	t.Logf("top is: %v", stack.Top())

	stack.Pop()
	t.Logf("top is: %v", stack.Top())

	stack.Pop()
	t.Logf("top is: %v", stack.Top())
	stack.Pop()
	t.Logf("top is: %v", stack.Top())
	stack.Pop()
	t.Logf("top is: %v", stack.Top())
}
