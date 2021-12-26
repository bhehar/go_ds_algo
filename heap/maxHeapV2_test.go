package heap

import (
	"testing"
	"reflect"
)

/*
	Tests below are by no means extensive. I wrote a couple tests
	simply to play around and get a taste of testing in GO.
*/

func TestInsertEmpty(t *testing.T) {
	maxHeap := &MaxHeap{}
	maxHeap.Insert(4)
	if maxHeap.Heap[0] != 4 && maxHeap.Count != 1 {
		t.Errorf("ERROR: got count: %v want count: %v", maxHeap.Count, 1)
		t.Errorf("ERROR: got heap: %v want heap: %v", maxHeap.Heap, []int{4})
	}
}

func TestRemoveEmpty(t *testing.T) {
	maxHeap := &MaxHeap{}
	val := maxHeap.Remove()
	// just checking to see if error is thrown
	if val != -1 {
		t.Error("ERROR: calling remove on an empty heap failed!")
	}
}

func TestInsert(t *testing.T) {
	maxHeap := &MaxHeap{}
	maxHeap.Insert(8)
	maxHeap.Insert(4)
	maxHeap.Insert(19)
	maxHeap.Insert(13)
	maxHeap.Insert(16)
	wantSlice := []int{19, 16, 14, 8, 4}
	if !reflect.DeepEqual(wantSlice, maxHeap.Heap) && maxHeap.Count != 5 {
		t.Errorf("ERROR: got count: %v want count: %v", maxHeap.Count, 5)
		t.Errorf("ERROR: got heap: %v want heap: %v", maxHeap.Heap, wantSlice)
	}
}

func TestRemove(t *testing.T) {
	maxHeap := &MaxHeap{}
	maxHeap.Insert(8)
	maxHeap.Insert(4)
	maxHeap.Insert(19)
	maxHeap.Insert(13)
	maxHeap.Insert(16)
	maxHeap.Remove()
	wantSlice := []int{16, 14, 8, 4}
	if !reflect.DeepEqual(wantSlice, maxHeap.Heap) && maxHeap.Count != 4 {
		t.Errorf("ERROR: got count: %v want count: %v", maxHeap.Count, 4)
		t.Errorf("ERROR: got heap: %v want heap: %v", maxHeap.Heap, wantSlice)
	}
}

