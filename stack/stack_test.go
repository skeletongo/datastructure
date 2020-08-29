package stack_test

import (
	"dataStructure/stack"
	"fmt"
	"testing"
)

func Example() {
	s := stack.NewArrayStack()
	fmt.Println(s)
	for i := 0; i < 3; i++ {
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

func BenchmarkArrayStack_Push(b *testing.B) {
	s := stack.NewArrayStack()
	for i := 0; i < b.N; i++ {
		s.Push(nil)
	}
	// BenchmarkArrayStack_Push-8      16476409                61.9 ns/op
}

func BenchmarkListStack_Push(b *testing.B) {
	s := stack.NewListStack()
	for i := 0; i < b.N; i++ {
		s.Push(nil)
	}
	// BenchmarkListStack_Push-8        9864056               115 ns/op
}

func BenchmarkNewArrayStack(b *testing.B) {
	b.StopTimer()
	s := stack.NewArrayStack()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Push(nil)
		s.Pop()
	}
	// BenchmarkNewArrayStack-8        355964994                3.34 ns/op
}

func BenchmarkNewListStack(b *testing.B) {
	b.StopTimer()
	s := stack.NewListStack()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Push(nil)
		s.Pop()
	}
	// BenchmarkNewListStack-8         20053474                51.0 ns/op
}
