package heap

import "fmt"

type MaxHeap struct {
	items []int
}

func (h *MaxHeap) Insert(key int) {
	h.items = append(h.items, key)
	h.maxHeapifyUp(len(h.items) - 1)
}

func (h *MaxHeap) Extract() int {
	if len(h.items) == 0 {
		fmt.Println("Can not extract from an empty heap")
		return -1
	}

	head := h.items[0]

	if len(h.items) == 1 {
		h.items = []int{}
		return head
	}
	lastIndex := len(h.items) - 1
	h.items[0] = h.items[lastIndex]
	h.items = h.items[:lastIndex]

	h.maxHeapifyDown(0)

	return head
}

func (h *MaxHeap) maxHeapifyUp(i int) {
	index := i
	parentIndex := getParentIndex(index)

	for h.items[parentIndex] < h.items[index] {
		h.swap(parentIndex, index)
		index = parentIndex
		parentIndex = getParentIndex(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	leftIndex := getLeftChildIndex(index)
	rightIndex := getRightChildIndex(index)
	childToCompare := 0
	lastIndex := len(h.items) - 1

	for leftIndex <= lastIndex {
		// left child is the only child
		if lastIndex == leftIndex {
			childToCompare = leftIndex
		} else if h.items[leftIndex] > h.items[rightIndex] { // left child is larger
			childToCompare = leftIndex
		} else { // right child is larger
			childToCompare = rightIndex
		}

		if h.items[index] < h.items[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			leftIndex = getLeftChildIndex(index)
			rightIndex = getRightChildIndex(index)
		} else {
			return
		}
	}

}

func (h *MaxHeap) swap(i1, i2 int) {
	h.items[i1], h.items[i2] = h.items[i2], h.items[i1]
}

func getParentIndex(i int) int {
	return (i - 1) / 2
}

func getLeftChildIndex(i int) int {
	return (2 * i) + 1
}

func getRightChildIndex(i int) int {
	return (2 * i) + 2
}
