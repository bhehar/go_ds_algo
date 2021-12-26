package heap

import (
	"fmt"
)

/*
	Heaps are 1 indexed so everytime we access the Arr, we subtract 1
*/
// insert adds an element to the end of the slice and calls heapify up
func Insert(arr []int, val int) []int {
	arr = append(arr, val)
	heapifyUp(arr)
	return arr
} 

// remove deletes an element from the root and calls heapify down
func remove(arr []int) []int {
	// did not have this during first implementation
	// lead to inddex out of range if slice was empty
	if len(arr) == 0 {
		fmt.Println("Cannot remove from an empty heap!")
		return arr
	} 
	swap(arr, 1, len(arr))
	arr = arr[:len(arr) - 1]
	heapifyDown(arr)
	return arr
}

// heapifyUp swaps last element upwards until correct position is found
func heapifyUp(arr []int) {
	targetIndx := len(arr)
	parentIndx := targetIndx / 2
	for arr[parentIndx - 1] < arr[targetIndx - 1] {
		swap(arr, targetIndx, parentIndx)
		targetIndx = parentIndx
		parentIndx = targetIndx / 2
		if parentIndx == 0 {
			return
		}
	}
}

// heapifyDown swapps root element downwards until correct position is found
// implementation has a bug. doing heapify down with only 1 element in slice leads to errors
func heapifyDown(arr []int) {
	parentIndx := 1
	leftChildIndx := parentIndx * 2
	rightChildIndx := parentIndx * 2 + 1
	maxChildIndx := max(arr, leftChildIndx, rightChildIndx)
	for( arr[parentIndx - 1] < arr[maxChildIndx - 1]) {
		swap(arr, parentIndx, maxChildIndx)
		parentIndx = maxChildIndx
		leftChildIndx = parentIndx * 2
		rightChildIndx = parentIndx * 2 + 1
		// we've reached the last node
		if parentIndx == len(arr) {
			return
		}
	}
}

// max returns max element given an slice and two indices
func max(arr []int, indxA, indxB int) int {
	if arr[indxB - 1] > arr[indxA - 1] {
		return indxB
	}
	return indxA
}

// swap exchnages elements at given indexes A & B in slice
func swap(arr []int, indxA, indxB int) {
	if indxA > len(arr) || indxB > len(arr) {
		fmt.Println("ERROR! Swap indexes must be less than arr len")
		return
	} 
	arr[indxA-1], arr[indxB-1] = arr[indxB-1], arr[indxA-1]
}