package linkedList

import (
	"fmt"
	"testing"
)

func TestLinkedListQueue(t *testing.T) {
	queue := new(Queue)
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

	fmt.Println("enqueue 5")
	queue.Enqueue(5)
	fmt.Println(queue)

	fmt.Println("get front:",queue.GetFront())
	fmt.Println(queue)
}
