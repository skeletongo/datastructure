package dataStructure

import (
	"container/ring"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	a := make([]int, 0, 5)
	//a = a[:cap(a)]
	fmt.Println(a)
}

func Test2(t *testing.T) {
	r := ring.New(0)
	r.Next()
}
