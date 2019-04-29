package array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	arr := NewArray(2)
	for i:=0;i<10;i++ {
		arr.AddLast(i)
		fmt.Println(arr)
	}

	arr = NewArray(2)
	arr.AddFirst(3)
	arr.AddLast(2)
	arr.AddLast(1)
	arr.AddFirst(4)
	fmt.Println(arr)
	arr.Insert(2,5)
	fmt.Println(arr)

	arr.Remove(2)
	fmt.Println(arr)

	arr.RemoveFirst()
	fmt.Println(arr)

	arr.RemoveLast()
	fmt.Println(arr)

	arr.Remove(0)
	fmt.Println(arr)

	arr.RemoveFirst()
	fmt.Println(arr)

	arr.AddFirst(1)
	fmt.Println(arr)

	arr.RemoveFirst()
	fmt.Println(arr)
}
