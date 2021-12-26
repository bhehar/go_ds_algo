package main

import (
	"fmt"
)

func sliceSandbox() {
	//play around here
}

// this is my initial implementation
func sliceInsertA( inputArr []string, index int, insertStr string) []string{
	elementsBeforeIndex := make([]string, index+1);
	copy(elementsBeforeIndex, inputArr[:index])
	elementsBeforeIndex[index] = insertStr;
	allElements := append(elementsBeforeIndex, inputArr[index:]...)
	return allElements;
}

// implementation I found online, it is similar
// to mine but concise
func sliceInsertB( inputArr []string, index int, insertStr string) []string {
	//inserting at the end or len(arr) == 0
	if(len(inputArr) == index) { 
		return append(inputArr, insertStr)
	}
	// we user :index+1 in the first arg to append bc we need to
	// make space for the new element
	inputArr = append(inputArr[:index+1], inputArr[index:]...)
	inputArr[index] = insertStr
	return inputArr
}

// prints slice with printStr, length, and capacity
func custPrint(desc string, slice []string) {
	printStr := desc + ": %#v Len: %v Cap %v\n"
	fmt.Printf(printStr, slice, len(slice), cap(slice))
}


// a little test for somthing we need to calculate permutations
func sliceExp1() {
	iniSlice := []string{"a", "b", "c"}
	for i := 0; i <= len(iniSlice); i++ {
		updatedSlice := sliceInsertB(iniSlice, i, "d")
		fmt.Println("Curr Index:", i)
		custPrint("d inserted into iniSlice", updatedSlice)
	}
}

// testing if modifying the resulting array of an append
// modified the original array
func sliceExp2() {
	iniSlice := []string{"a", "b", "c"}
	// newSlice := iniSlice[:4]
	newSlice := append(iniSlice, "d")
	newSlice[0] = "z"
	custPrint("iniSlice after append and update", iniSlice)
	custPrint("new slice", newSlice)
}

// playing around with copy()
func sliceExp3() {
	arr1 := []string{"a", "b", "c"}
	arr2 := make([]string, 3, 5)
	// custPrint("arr1 ini", arr1)
	// custPrint("arr2 ini", arr2)

	// let's play with copy
	// in order to copy the receiving array must have
	// the proper length, capacity doesn't matter
	copy(arr2, arr1)
	// custPrint("arr1 copied -> arr2 w/ copy()", arr2)
	// copy onto itself
	copy(arr2, arr2[1:])
	custPrint("arr2[1:] copied onto intself", arr2)
}

// testing slice creation from another slice
// points to same underlying memory
func sliceExp4() {
	arr1 := []string{"a", "b", "c"}
	arr2 := arr1[:0]
	arr2 = append(arr2, "d")
	custPrint("arr1", arr1)
	custPrint("arr2", arr2)
}