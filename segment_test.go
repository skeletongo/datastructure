package dataStructure

import (
	"testing"
)

func TestSegmentTree(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	test := [][]int{
		{0, 2, 1},
		{2, 4, 0},
		{3, 5, -4},
	}

	s := &SegmentTree{}
	s.NewTree(nums, func(a, b int) int {
		return a + b
	})
	for _, v := range test {
		if s.Query(v[0], v[1]) != v[2] {
			t.Error("线段树查询错误")
		}
	}

	s.Set(0, -1)
	s.Set(2, 5)
	s.Set(5, -3)

	test = [][]int{
		{0, 2, 4},
		{2, 4, 2},
		{3, 5, -6},
	}
	for _, v := range test {
		if s.Query(v[0], v[1]) != v[2] {
			t.Error("线段树修改错误")
		}
	}
}
