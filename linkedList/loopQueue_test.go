package linkedList

import (
	"fmt"
	"testing"
)

func TestLoopQueue(t *testing.T) {
	queue := new(LoopQueue)
	for i:=0;i<5;i++{
		queue.Enqueue(i)
		fmt.Println(queue)
	}

	fmt.Println("dequeue:",queue.Dequeue())
	fmt.Println(queue)

	fmt.Println("dequeue:",queue.Dequeue())
	fmt.Println(queue)

	fmt.Println("dequeue:",queue.Dequeue())
	fmt.Println(queue)

	fmt.Println("enqueue 6")
	queue.Enqueue(6)
	fmt.Println(queue)

	fmt.Println("get front:",queue.GetFront())
	fmt.Println(queue)
}
