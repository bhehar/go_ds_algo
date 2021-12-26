package main

func BinarySearchIter(arr []int, n, target int) int {
	low := 0
	high := n
	for low <= high {
		mid := (low + high) / 2
		if target == arr[mid] {
			return mid
		} else if target < arr[mid]{
			high = mid
		} else {
			low = mid
		}
	}
	return -1
}

func BinarySearchRecur(arr []int, low, high, target int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if target < arr[mid] {
		return BinarySearchRecur(arr, low, mid - 1, target)
	} else if target > arr[mid] {
		return BinarySearchRecur(arr, mid + 1, high, target)
	}
	return mid
}
