package array

import (
	"fmt"
	"testing"
)

func TestArrayQueue(t *testing.T) {
	queue := new(Queue)
	queue.Init()
	for i := 0; i < 5; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)
	}

	for i := 0; i < 5; i++ {
		fmt.Println("front", queue.GetFront())
		fmt.Println("dequeue", queue.Dequeue())
		fmt.Println(queue)
	}
}
