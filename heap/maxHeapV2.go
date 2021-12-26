package heap

import (
	"fmt"
)

/*
	Imporvements were made after watchign video by Junmin Lee on YouTube
	Link: https://www.youtube.com/watch?v=3DYIgTC4T1o
*/

type MaxHeap struct {
	Heap  []int
	Count int
}

func (h *MaxHeap) Insert(val int) {
	h.Heap = append(h.Heap, val)
	h.Count++
	h.heapifyUp()
}

func (h *MaxHeap) Remove() int {
	if h.Count == 0 {
		// fmt.Println("ERROR! Cannot remove from an empty heap!")
		return -1
	}
	// retrive root for return
	elm := h.Heap[0]
	// move last node into root
	h.swap(0, h.Count-1)
	// decrement count
	h.Count--
	// reslice h.heap
	h.Heap = h.Heap[:h.Count]
	h.heapifyDown()
	return elm
}

func (h *MaxHeap) heapifyDown() {
	parent := 0
	// as long as we're not looking at the last index
	for parent < h.Count-1 {
		// child to compare against
		childToComp := h.childToComp(parent)
		if h.Heap[parent] > h.Heap[childToComp] {
			//if parent is bigger than larger child, we're done
			return
		}
		// swap
		h.swap(parent, childToComp)
		// update parent to new index
		parent = childToComp
	}
}

func (h *MaxHeap) heapifyUp() {
	target := h.Count - 1
	parent := parentIndx(target)
	for h.Heap[target] > h.Heap[parent] {
		h.swap(target, parent)
		// if we're at root node, we're done
		if target == 0 {
			return
		}
		target = parent
		parent = parentIndx(target)
	}
}

// childToComp returns the larger child index given a parent index
func (h *MaxHeap) childToComp(parent int) int {
	lc := leftChildIndx(parent)
	rc := rightChildIndx(parent)
	childToComp := -1
	if lc == h.Count-1 {
		// if leftChild is last element
		childToComp = lc
	} else if h.Heap[lc] > h.Heap[rc] {
		// if leftChild > right(child) then leftChild
		childToComp = lc
	} else {
		// right child
		childToComp = rc
	}
	return childToComp
}

func (h *MaxHeap) swap(a, b int) {
	h.Heap[a], h.Heap[b] = h.Heap[b], h.Heap[a]
}

// childToCompare returns the index of child we want to compare with

// leftChild return indx of left child in heap slice/arr
func leftChildIndx(i int) int {
	return (i * 2) + 1
}

// parent returns the index of a parent node given child node index
func parentIndx(i int) int {
	return (i - 1) / 2
}

// rightChild return indx of right child in heap slice/arr
func rightChildIndx(i int) int {
	return (i * 2) + 2
}

// -------- funcitons to test this implementation --------
// insertEmpty attempt to insert a node into an empty heap
func V2DisplayInsert() {
	mh := MaxHeap{}
	toInsert := []int{4, 8, 19, 16, 13}
	fmt.Printf("Initial Heap: %v\n", mh.Heap)
	for _, item := range toInsert {
		mh.Insert(item)
		fmt.Printf("\nInserted: %v\nHeap: %v\n", item, mh.Heap)
	}
}

func V2DisplayRemove() {
	mh := MaxHeap{
		Heap:  []int{19, 16, 13, 8, 4},
		Count: 5,
	}
	fmt.Printf("Initial Heap: %v\n", mh.Heap)
	for i := 0; i < 5; i++ {
		r := mh.Remove()
		fmt.Printf("\nRemoved: %v\nHeap: %v\n", r, mh.Heap)
	}
}
