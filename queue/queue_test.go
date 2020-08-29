package queue

import (
	"dataStructure/queue/array"
	"dataStructure/queue/ring"
	"errors"
	"math/rand"
	"testing"
)

func testQueue(q IQueue) error {
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
	for e := q.Dequeue(); e != nil; e = q.Dequeue() {
		res = append(res, e.(int))
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
	if err := testQueue(array.New()); err != nil {
		t.Error(err)
	}
}

func TestRingQueue(t *testing.T) {
	if err := testQueue(ring.New()); err != nil {
		t.Error(err)
	}
}

func BenchmarkArrayQueueEnqueue(b *testing.B) {
	q := array.New()
	for i := 0; i < b.N; i++ {
		q.Enqueue(nil)
	}
	// BenchmarkArrayQueueEnqueue-20           20000000               104 ns/op
}

func BenchmarkRingQueueEnqueue(b *testing.B) {
	q := ring.New()
	for i := 0; i < b.N; i++ {
		q.Enqueue(nil)
	}
	// BenchmarkRingQueueEnqueue-20            20000000                88.2 ns/op
}

const EnqueueNum = 1000000

func BenchmarkArrayQueueDequeue(b *testing.B) {
	b.StopTimer()
	q := array.New()
	for i := 0; i < EnqueueNum; i++ {
		q.Enqueue(nil)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if i > EnqueueNum {
			break
		}
		q.Dequeue()
	}
	// BenchmarkArrayQueueDequeue-20               1000           1415131 ns/op
}

func BenchmarkRingQueueDequeue(b *testing.B) {
	b.StopTimer()
	q := ring.New()
	for i := 0; i < EnqueueNum; i++ {
		q.Enqueue(nil)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if i > EnqueueNum {
			break
		}
		q.Dequeue()
	}
	// BenchmarkRingQueueDequeue-20            2000000000               0.02 ns/op
}
