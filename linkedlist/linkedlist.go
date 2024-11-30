package linkedlist

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (linkedList *LinkedList) Prepend(value *Node) {
	second := linkedList.head
	linkedList.head = value
	linkedList.head.next = second

	linkedList.length++
}

func (linkedList *LinkedList) Append(value *Node) {
	if linkedList.head == nil {
		linkedList.head = value
		linkedList.length++
	} else {
		current := linkedList.head
		for current.next != nil {
			current = current.next
		}
		current.next = value
		linkedList.length++
	}
}

func (linkedList *LinkedList) ReplaceNode(index int, value *Node) int {
	if linkedList.head == nil {
		return -1
	}
	current := linkedList.head
	for i := 0; i < index; i++ {
		if current.next == nil {
			return -1
		}
		current = current.next
	}
	current.value = value.value
	return 1
}

func (linkedList *LinkedList) DeleteWithValue(value int) int {
	if linkedList.head == nil {
		return -1
	}
	if linkedList.head.value == value {
		linkedList.head = linkedList.head.next
		linkedList.length--
		return 1
	}
	current := linkedList.head
	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			linkedList.length--
			return 1
		}
		current = current.next
	}
	return -1
}
