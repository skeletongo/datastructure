package array_test

import (
	"dataStructure/array"
	"fmt"
)

func ExampleNew() {
	arr := array.New()
	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr)

	arr.RemoveFirst()
	arr.RemoveLast()
	fmt.Println(arr)

	arr.AddLast(9)
	arr.AddFirst(0)
	fmt.Println(arr)

	arr.Insert(1, 10)
	fmt.Println(arr)

	arr.Remove(1)
	fmt.Println(arr)

	arr.Set(0, -1)
	fmt.Println(arr)

	arr.Swap(0, 1)
	fmt.Println(arr)

	fmt.Println(arr.Contains(5))

	fmt.Println(arr.Get(1))
	// Output:
	// len:10 cap:16 array: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	// len:8 cap:16 array: [1, 2, 3, 4, 5, 6, 7, 8]
	// len:10 cap:16 array: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	// len:11 cap:16 array: [0, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	// len:10 cap:16 array: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	// len:10 cap:16 array: [-1, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	// len:10 cap:16 array: [1, -1, 2, 3, 4, 5, 6, 7, 8, 9]
	// true
	// -1
}

func ExampleArray_ContainsFunc() {
	type O struct {
		a string
		b int
	}

	o1 := &O{
		a: "a",
		b: 1,
	}

	o2 := &O{
		a: "b",
		b: 1,
	}

	o3 := &O{
		a: "a",
		b: 1,
	}

	arr := array.New()
	arr.AddLast(o1)
	fmt.Println(arr.Contains(o2))
	fmt.Println(arr.Contains(o3))

	f := func(i, j interface{}) bool {
		o1 := i.(*O)
		o2 := j.(*O)
		return o1.b == o2.b
	}
	fmt.Println(arr.ContainsFunc(o2, f))
	fmt.Println(arr.ContainsFunc(o3, f))
	// Output:
	// false
	// true
	// true
	// true
}
