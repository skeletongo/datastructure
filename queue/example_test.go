package queue_test

import (
	"fmt"

	"github.com/skeletongo/dataStructure/queue"
)

func ExampleArrayQueue() {
	q := queue.NewArrayQueue()
	fmt.Println(q.Len() == 0)
	fmt.Println(q.Cap())
	for i := 0; i < 4; i++ {
		q.Enqueue(i)
	}
	fmt.Println(q.Len() == 0)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Cap())
	fmt.Println(q.Len() == 0)
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
	q.Enqueue(6)
	fmt.Println(q)

	// Output:
	// true
	// 1
	// false
	// len: 4 queue: front [0, 1, 2, 3]
	// 0
	// len: 3 queue: front [1, 2, 3]
	// 1
	// 2
	// 3
	// <nil>
	// 2
	// true
	// len: 5 queue: front [0, 1, 2, 3, 4]
	// len: 4 queue: front [1, 2, 3, 4]
	// len: 5 queue: front [1, 2, 3, 4, 6]
}

func ExampleListQueue() {
	q := queue.NewListQueue()
	fmt.Println(q.Len() == 0)
	for i := 0; i < 4; i++ {
		q.Enqueue(i)
	}
	fmt.Println(q.Len() == 0)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Len() == 0)
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
	q.Enqueue(6)
	fmt.Println(q)

	// Output:
	// true
	// false
	// len: 4 queue: front [0, 1, 2, 3]
	// 0
	// len: 3 queue: front [1, 2, 3]
	// 1
	// 2
	// 3
	// <nil>
	// true
	// len: 5 queue: front [0, 1, 2, 3, 4]
	// len: 4 queue: front [1, 2, 3, 4]
	// len: 5 queue: front [1, 2, 3, 4, 6]
}

func ExampleRingQueue() {
	q := queue.NewRingQueue()
	fmt.Println(q.Len() == 0)
	for i := 0; i < 4; i++ {
		q.Enqueue(i)
	}
	fmt.Println(q.Len() == 0)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Len() == 0)
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
	q.Enqueue(6)
	fmt.Println(q)

	// Output:
	// true
	// false
	// len: 4 queue: front [0, 1, 2, 3]
	// 0
	// len: 3 queue: front [1, 2, 3]
	// 1
	// 2
	// 3
	// <nil>
	// true
	// len: 5 queue: front [0, 1, 2, 3, 4]
	// len: 4 queue: front [1, 2, 3, 4]
	// len: 5 queue: front [1, 2, 3, 4, 6]
}
