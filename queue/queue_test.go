package queue

import (
	"errors"
	"math/rand"
	"testing"
)

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

func BenchmarkArrayQueue(b *testing.B) {
	b.StopTimer()
	q := NewArrayQueue()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		q.Enqueue(nil)
	}
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
	// BenchmarkLoopArrayQueue-8       19110135                74.1 ns/op
}

func BenchmarkListQueue(b *testing.B) {
	b.StopTimer()
	q := NewListQueue()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		q.Enqueue(nil)
	}
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
	// BenchmarkLineListQueue-8         6716150               154 ns/op
}

func BenchmarkRingQueue(b *testing.B) {
	b.StopTimer()
	q := NewRingQueue()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		q.Enqueue(nil)
	}
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
	// BenchmarkLoopRingQueue-8        10108131               108 ns/op
}
