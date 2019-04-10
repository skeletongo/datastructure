package dataStructure

import (
	"math/rand"
	"testing"
	"time"
)

func TestMaxHeap(t *testing.T) {
	maxHeap := MaxHeap{}
	n := 1000
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		maxHeap.Add(r.Intn(10000))
	}

	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, maxHeap.ExtractMax())
	}

	for i := 1; i < n; i++ {
		if arr[i-1] < arr[i] {
			t.Error("堆错误")
		}
	}
}

func TestMaxHeap_Heapify(t *testing.T) {
	maxHeap := MaxHeap{}
	n := 1000
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, r.Intn(10000))
	}

	maxHeap.Heapify(arr)

	for i := 0; i < n; i++ {
		arr[i] = maxHeap.ExtractMax()
	}

	for i := 1; i < n; i++ {
		if arr[i-1] < arr[i] {
			t.Error("堆错误")
		}
	}
}
