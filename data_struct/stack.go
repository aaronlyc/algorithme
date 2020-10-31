package data_struct

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack
func NewListStack() *Stack {
	return &Stack{nil, 0}
}

// Return the number of items in the stack
func (s *Stack) Len() int {
	return s.length
}

// View the top item on the stack
func (s *Stack) Top() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

// Pop the top item of the stack and return it
func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

// Push a value onto the top of the stack
func (s *Stack) Push(value interface{}) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}

/////////////////////////////////// 基于slice创建的stack
// ItemStack the stack of Items
type ItemStack struct {
	items []int
	i     int // 表示栈顶在items中的位置
}

// New creates a new ItemStack
// 可固定长度
func NewSliceStack(len int) *ItemStack {
	s := &ItemStack{
		items: make([]int, len, len),
		i:     -1,
	}
	return s
}

func (s *ItemStack) Items() []int {
	return s.items[:s.i+1]
}

// Pirnt prints all the elements
func (s *ItemStack) Top() interface{} {
	if s.Empty() {
		return nil
	}
	return s.items[s.i]
}

// Empty is ask the stack is empty or not
func (s *ItemStack) Empty() bool {
	return s.i == -1
}

// Push adds an Item to the top of the stack
func (s *ItemStack) Push(t int) {
	s.i++
	if len(s.items) > s.i {
		s.items[s.i] = t
	} else {
		s.items = append(s.items, t)
	}
}

// Pop removes an Item from the top of the stack
func (s *ItemStack) Pop() {
	if s.Empty() {
		return
	}
	s.i--
}
