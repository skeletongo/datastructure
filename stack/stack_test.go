package stack_test

import (
	"dataStructure/stack"
	"fmt"
)

func ExampleNew() {
	s := stack.New()
	fmt.Println(s)
	for i :=0; i < 3; i++ {
		s.Push(i)
	}
	fmt.Println(s)
	for i := 0; i < 2; i++ {
		fmt.Println(s.Pop())
	}
	fmt.Println(s)
	s.Push(3)
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	// Output:
	// len: 0 stack: [] top
	// len: 3 stack: [0, 1, 2] top
	// 2
	// 1
	// len: 1 stack: [0] top
	// len: 2 stack: [0, 3] top
	// 3
	// 0
	// <nil>
}
