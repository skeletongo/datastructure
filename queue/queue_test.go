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
