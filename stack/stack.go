package stack

type Stack struct {
	items []int
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.items) == 0
}

func (stack *Stack) GetLength() int {
	return len(stack.items)
}

func (stack *Stack) Push(item int) {
	stack.items = append(stack.items, item)
}

func (stack *Stack) Pop() int {
	if stack.IsEmpty() {
		return -1
	}
	item := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return item
}

func (stack *Stack) Peek() int {
	if stack.IsEmpty() {
		return -1
	}
	return stack.items[len(stack.items)-1]
}
