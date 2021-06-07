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

	fmt.Println(arr.Contains(5, func(a, b interface{}) bool {
		return a.(int) == b.(int)
	}))

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
