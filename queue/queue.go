package queue

type Queue struct {
	items []int
}

/*
 returns the length of the queue
*/
func (queue *Queue) GetLength() int {
	return len(queue.items)
}

/*
 returns true if the queue is empty
*/
func (queue *Queue) IsEmpty() bool {
	return len(queue.items) == 0
}

// Enqueue adds an item to the end of the queue.
func (queue *Queue) Enqueue(item int) {
	queue.items = append(queue.items, item)
}

/*
 Dequeue removes and returns the first item in the queue
*/
func (queue *Queue) Dequeue() int {
	if (queue.IsEmpty()) {
		return -1
	}
	item := queue.items[0]
	queue.items = queue.items[1:]
	return item
}

/*
 Peek returns the first item in the queue without removing it
*/
func (queue *Queue) Peek() int {
	if (queue.IsEmpty()) {
		return -1
	}
	return queue.items[0]
}
