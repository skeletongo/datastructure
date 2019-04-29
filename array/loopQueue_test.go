package array

import (
	"fmt"
	"testing"
)

func TestLoopQueue(t *testing.T) {
	queue := new(LoopQueue)
	queue.Init()
	for i:=0;i<5;i++ {
		queue.Enqueue(i)
		fmt.Println(queue)
	}

	for i:=0;i<5;i++ {
		queue.Dequeue()
		fmt.Println(queue)
	}

	for i:=0;i<5;i++ {
		queue.Enqueue(i)
		fmt.Println(queue)
	}

	for i:=0;i<5;i++ {
		queue.Dequeue()
		fmt.Println(queue)
	}

}
