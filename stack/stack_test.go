package stack_test

import (
	"fmt"
	"testing"

	"github.com/skeletongo/dataStructure/stack"
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

func BenchmarkStack_Push(b *testing.B) {
	n := 1000

	newFunc := func() stack.Stack {
		return stack.NewArrayStack()
	}

	f := func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			s := newFunc()
			b.StartTimer()
			for j := 0; j < n; j++ {
				s.Push(nil)
			}
		}
	}

	b.Run("ArrayStackPush", f)

	newFunc = func() stack.Stack {
		return stack.NewListStack()
	}
	b.Run("ListStackPush", f)

	//goos: windows
	//goarch: amd64
	//pkg: dataStructure/stack
	//cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
	//BenchmarkStack_Push
	//BenchmarkStack_Push/ArrayStackPush
	//BenchmarkStack_Push/ArrayStackPush-8         	  109525	     11081 ns/op
	//BenchmarkStack_Push/ListStackPush
	//BenchmarkStack_Push/ListStackPush-8          	   29740	     37685 ns/op
}

func BenchmarkStack_Pop(b *testing.B) {
	n := 1000

	newFunc := func() stack.Stack {
		return stack.NewArrayStack()
	}

	f := func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			s := newFunc()
			for j := 0; j < n; j++ {
				s.Push(nil)
			}
			b.StartTimer()
			for j := 0; j < n; j++ {
				s.Pop()
			}
		}
	}

	b.Run("ArrayStackPop", f)

	newFunc = func() stack.Stack {
		return stack.NewListStack()
	}
	b.Run("ListStackPop", f)

	//goos: windows
	//goarch: amd64
	//pkg: dataStructure/stack
	//cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
	//BenchmarkStack_Pop
	//BenchmarkStack_Pop/ArrayStackPop
	//BenchmarkStack_Pop/ArrayStackPop-8         	  539031	      2451 ns/op
	//BenchmarkStack_Pop/ListStackPop
	//BenchmarkStack_Pop/ListStackPop-8          	  215617	      5293 ns/op
}
