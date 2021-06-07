package heap

import (
	"math/rand"
	"testing"
	"time"
)

func TestHeap(t *testing.T) {
	maxHeap := New(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	n := 1000
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		maxHeap.Add(r.Intn(10000))
	}

	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, maxHeap.ExtractMax().(int))
	}

	for i := 1; i < n; i++ {
		if arr[i-1] < arr[i] {
			t.Error("堆错误")
		}
	}
}

func TestHeap_Heapify(t *testing.T) {
	maxHeap := New(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	n := 1000
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var arr []interface{}
	for i := 0; i < n; i++ {
		arr = append(arr, r.Intn(10000))
	}

	maxHeap.Heapify(arr)

	for i := 0; i < n; i++ {
		arr[i] = maxHeap.ExtractMax()
	}

	for i := 1; i < n; i++ {
		if arr[i-1].(int) < arr[i].(int) {
			t.Error("堆错误")
		}
	}
}
