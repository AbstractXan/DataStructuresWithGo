package stack

import l "list"

// Stack FIFO stack with push and pop
// Push | Pop | Top
type Stack struct {
	list l.List
}

// Push pushes value on stack
func (stack *Stack) Push(newVal int) {
	stack.list.Insert(newVal)
}

// Top returns value at top
func (stack *Stack) Top() int {
	return stack.list.GetAt(stack.list.Length())
}

// Pop returns value at top and removes it from the stack
func (stack *Stack) Pop() int {
	popped := stack.Top()
	stack.list.DeleteLast()
	return popped
}

// Length returns the number of elements in the stack
func (stack *Stack) Length() int {
	return stack.list.Length()
}
