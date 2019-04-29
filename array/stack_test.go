package array

import (
	"fmt"
	"testing"
)

func TestArrayStack(t *testing.T) {
	stack := new(Stack)
	stack.Init()
	for i:=0;i<5;i++ {
		stack.Push(i)
		fmt.Println(stack)
	}

	for i:=0;i<5;i++ {
		fmt.Println("peek",stack.Peek())
		fmt.Println("pop",stack.Pop())
		fmt.Println(stack)
	}
}
