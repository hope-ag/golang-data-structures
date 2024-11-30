package heap

import "fmt"

func TryMaxHeap() {
	maxHeap := MaxHeap{}

	fmt.Println(maxHeap)

	opts := []int{10, 20, 11, 54, 0, 90, 1, 6, 30}

	for _, num := range opts {
		maxHeap.Insert(num)
		fmt.Println(maxHeap)
	}

	fmt.Println("Extracting top value")
	
	val := maxHeap.Extract()

	fmt.Println(val)

	fmt.Println(maxHeap)

}