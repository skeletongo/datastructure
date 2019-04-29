package linkedList

import (
	"fmt"
	"testing"
)

func TestLinkedListStack(t *testing.T) {
	stack := new(Stack)
	for i:=0;i<5;i++ {
		stack.Push(i)
		fmt.Println(stack)
	}

	fmt.Println("pop:",stack.Pop())
	fmt.Println(stack)

	fmt.Println("push 6")
	stack.Push(6)
	fmt.Println(stack)

	for i:=0;i<3;i++ {
		stack.Pop()
		fmt.Println(stack)
	}

	fmt.Println(stack.Peek())
	fmt.Println(stack)
}
