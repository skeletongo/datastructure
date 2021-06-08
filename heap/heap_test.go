package heap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestHeap(t *testing.T) {
	maxHeap := NewArrayHeap(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	n := 1000
	for i := 0; i < n; i++ {
		maxHeap.Add(rand.Intn(10000))
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
	maxHeap := NewArrayHeap(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	n := 1000
	var arr []interface{}
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(10000))
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

func TestArrayHeap_String(t *testing.T) {
	maxHeap := NewArrayHeap(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	n := 20
	for i := 0; i < n; i++ {
		maxHeap.Add(rand.Intn(50))
	}
	fmt.Println(maxHeap)
}
