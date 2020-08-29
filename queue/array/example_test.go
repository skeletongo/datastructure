package array_test

import (
	"dataStructure/queue/array"
	"fmt"
)

func ExampleNew() {
	q := array.New()
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Cap())
	for i := 0; i < 4; i++ {
		q.Enqueue(i)
	}
	fmt.Println(q.IsEmpty())
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Cap())
	fmt.Println(q.IsEmpty())
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
	// 0
	// false
	// len: 4 queue: front [0, 1, 2, 3]
	// 0
	// len: 3 queue: front [1, 2, 3]
	// 1
	// 2
	// 3
	// <nil>
	// 4
	// true
	// len: 5 queue: front [0, 1, 2, 3, 4]
	// len: 4 queue: front [1, 2, 3, 4]
	// len: 5 queue: front [1, 2, 3, 4, 6]
}
