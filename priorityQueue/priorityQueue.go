package main

import "fmt"

// implemented with the heap algorithm

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Display() {
	fmt.Println(h.array)
}

// after insertion or removal we need to satify the heap property

// insert adds an element to the heap
func (h *MaxHeap) Insert(elt int) {
	h.array = append(h.array, elt)
	h.heapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func parent(i int) int {
	return (i - 1) / 2
}

func (h *MaxHeap) heapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) heapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { // left is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // left is the larger child
			childToCompare = l
		} else { // right is the larger child
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// extract returns the largest key and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	// replace first value with last value
	last := len(h.array) - 1
	h.array[0] = h.array[last]
	// remove last element
	h.array = h.array[:last]
	h.heapifyDown(0)
	return extracted
}

// extract a selected value from the priority queue
func (h *MaxHeap) ExtractValue(value int) int {
	return 1
}

func main() {
	fmt.Println("Heaps !")
	buildHeap := []int{10, 20, 30, 5, 6, 1}
	priorityQueue := MaxHeap{}
	for _, v := range buildHeap {
		priorityQueue.Insert(v)
	}
	priorityQueue.Extract()
	priorityQueue.Display()
}
