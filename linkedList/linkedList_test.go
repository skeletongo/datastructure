package linkedList

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := new(LinkedList)
	for i := 0; i < 5; i++ {
		list.AddLast(i)
		fmt.Println(list)
	}

	list = new(LinkedList)
	for i := 0; i < 5; i++ {
		list.AddFirst(i)
		fmt.Println(list)
	}

	fmt.Println("remove first:", list.RemoveFirst())
	fmt.Println(list)

	fmt.Println("remove last:", list.RemoveLast())
	fmt.Println(list)

	fmt.Println("remove 1:", list.Remove(1))
	fmt.Println(list)

	fmt.Println("get first:", list.GetFirst())
	fmt.Println(list)

	fmt.Println("get last:", list.GetLast())
	fmt.Println(list)

	fmt.Println("contains 4:",list.Contains(4))
	fmt.Println(list)

	fmt.Println("contains 3:",list.Contains(3))
	fmt.Println(list)

	fmt.Println("set 1 4")
	list.Set(1,4)
	fmt.Println(list)

	fmt.Println("remove 4")
	list.RemoveElement(4)
	fmt.Println(list)

	fmt.Println("remove 3")
	list.RemoveElement(3)
	fmt.Println(list)
}
