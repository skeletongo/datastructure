package queue

import (
	"errors"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func testQueue(q Queue) error {
	var data []int
	var res []int

	for i := 0; i < 1000000; i++ {
		if rand.Intn(2) == 0 {
			n := rand.Int()
			q.Enqueue(n)
			data = append(data, n)
		} else {
			n, ok := q.Dequeue().(int)
			if ok {
				res = append(res, n)
			}
		}
	}
	for v := q.Dequeue(); v != nil; v = q.Dequeue() {
		res = append(res, v.(int))
	}

	if len(data) != len(res) {
		return errors.New("1")
	}
	for i := 0; i < len(data); i++ {
		if data[i] != res[i] {
			return errors.New("2")
		}
	}
	return nil
}

func testHeapQueue(q *PriorityQueue) error {
	n := 100000
	var data []int
	for i := 0; i < n; i++ {
		v := rand.Intn(n)
		item := &Item{Value: v}
		q.Enqueue(item)
		if rand.Intn(2) == 0 {
			item.Value = rand.Intn(n)
			HeapUpdate(q, item)
		}
	}
	for v := q.Dequeue(); v != nil; v = q.Dequeue() {
		data = append(data, v.(*Item).Value.(int))
	}

	if len(data) != n {
		return errors.New("1")
	}
	for i := 1; i < n; i++ {
		if data[i-1] > data[i] {
			return errors.New("2")
		}
	}
	return nil
}

func TestArrayQueue(t *testing.T) {
	if err := testQueue(NewArrayQueue()); err != nil {
		t.Error(err)
	}
}

func TestListQueue(t *testing.T) {
	if err := testQueue(NewListQueue()); err != nil {
		t.Error(err)
	}
}

func TestLoopRingQueue(t *testing.T) {
	if err := testQueue(NewRingQueue()); err != nil {
		t.Error(err)
	}
}

func TestHeapQueue(t *testing.T) {
	if err := testHeapQueue(NewPriorityQueue(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})); err != nil {
		t.Error(err)
	}
}

// 优先队列入队测试
func BenchmarkHeapPush(b *testing.B) {
	n := 1000
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		q := NewPriorityQueue(func(a, b interface{}) int {
			return a.(int) - b.(int)
		})
		b.StartTimer()
		for j := 0; j < n; j++ {
			q.Enqueue(&Item{Value: rand.Int()})
		}
	}

	//goos: windows
	//goarch: amd64
	//pkg: dataStructure/queue
	//cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
	//BenchmarkHeapPush
	//BenchmarkHeapPush-8   	   12266	     95403 ns/op
}

// 优先队列出队测试
func BenchmarkHeapPop(b *testing.B) {
	n := 1000
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		q := NewPriorityQueue(func(a, b interface{}) int {
			return a.(int) - b.(int)
		})
		for j := 0; j < n; j++ {
			q.Enqueue(&Item{Value: rand.Int()})
		}
		b.StartTimer()
		for j := 0; j < n; j++ {
			q.Dequeue()
		}
	}

	//goos: windows
	//goarch: amd64
	//pkg: dataStructure/queue
	//cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
	//BenchmarkHeapPop
	//BenchmarkHeapPop-8   	    5824	    190853 ns/op
}

// 入队性能测试
func BenchmarkQueue_Enqueue(b *testing.B) {
	n := 1000
	newFunc := func() Queue {
		return NewArrayQueue()
	}

	f := func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			q := newFunc()
			b.StartTimer()
			for j := 0; j < n; j++ {
				q.Enqueue(nil)
			}
		}
	}

	b.Run("ArrayQueueEnqueue", f)

	newFunc = func() Queue {
		return NewListQueue()
	}
	b.Run("ListQueueEnqueue", f)

	newFunc = func() Queue {
		return NewRingQueue()
	}
	b.Run("RingQueueEnqueue", f)

	//goos: windows
	//goarch: amd64
	//pkg: dataStructure/queue
	//cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
	//BenchmarkQueue_Enqueue
	//BenchmarkQueue_Enqueue/ArrayQueueEnqueue
	//BenchmarkQueue_Enqueue/ArrayQueueEnqueue-8         	   41953	     29047 ns/op
	//BenchmarkQueue_Enqueue/ListQueueEnqueue
	//BenchmarkQueue_Enqueue/ListQueueEnqueue-8          	   31330	     40473 ns/op
	//BenchmarkQueue_Enqueue/RingQueueEnqueue
	//BenchmarkQueue_Enqueue/RingQueueEnqueue-8          	   34465	     34310 ns/op
}

// 出队性能测试
func BenchmarkQueue_Dequeue(b *testing.B) {
	n := 1000
	newFunc := func() Queue {
		return NewArrayQueue()
	}

	f := func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			q := newFunc()
			for j := 0; j < n; j++ {
				q.Enqueue(nil)
			}
			b.StartTimer()
			for j := 0; j < n; j++ {
				q.Dequeue()
			}
		}
	}

	b.Run("ArrayQueueDequeue", f)

	newFunc = func() Queue {
		return NewListQueue()
	}
	b.Run("ListQueueDequeue", f)

	newFunc = func() Queue {
		return NewRingQueue()
	}
	b.Run("RingQueueDequeue", f)

	//goos: windows
	//goarch: amd64
	//pkg: dataStructure/queue
	//cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
	//BenchmarkQueue_Dequeue
	//BenchmarkQueue_Dequeue/ArrayQueueDequeue
	//BenchmarkQueue_Dequeue/ArrayQueueDequeue-8         	   56192	     20558 ns/op
	//BenchmarkQueue_Dequeue/ListQueueDequeue
	//BenchmarkQueue_Dequeue/ListQueueDequeue-8          	  200180	      5204 ns/op
	//BenchmarkQueue_Dequeue/RingQueueDequeue
	//BenchmarkQueue_Dequeue/RingQueueDequeue-8          	  142177	      8465 ns/op
}

// 综合性能测试
func BenchmarkQueue(b *testing.B) {
	n := 1000000
	arr := rand.Perm(n)
	newFunc := func() Queue {
		return NewArrayQueue()
	}

	f := func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			q := newFunc()
			b.StartTimer()
			for j := 0; j < n; j++ {
				if arr[j]%2 == 0 {
					q.Enqueue(nil)
				} else {
					q.Dequeue()
				}
			}
		}
	}

	b.Run("ArrayQueue", f)

	newFunc = func() Queue {
		return NewListQueue()
	}
	b.Run("ListQueue", f)

	newFunc = func() Queue {
		return NewRingQueue()
	}
	b.Run("RingQueue", f)

	//goos: windows
	//goarch: amd64
	//pkg: dataStructure/queue
	//cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
	//BenchmarkQueue
	//BenchmarkQueue/ArrayQueue
	//BenchmarkQueue/ArrayQueue-8         	      69	  16000912 ns/op
	//BenchmarkQueue/ListQueue
	//BenchmarkQueue/ListQueue-8          	      44	  26271632 ns/op
	//BenchmarkQueue/RingQueue
	//BenchmarkQueue/RingQueue-8          	      48	  25762088 ns/op
}
