package heap

import "fmt"

// Heap creates a min-heap
type Heap struct {
	Array []int
	size  int
}

func (heap *Heap) Insert(data int) {
	heap.Array = append(heap.Array, data)
	heap.size++
}

func (heap *Heap) Print() {
	fmt.Println(heap.Array)
}
