package dataStructure

import (
	"fmt"
	"testing"
)

func Same(b Bst, f func(bst Bst) []int, res []int) bool {
	list := f(b)
	if len(list) != len(res) {
		return false
	}
	for k, v := range res {
		if list[k] != v {
			return false
		}
	}
	return true
}

func Pt(list []int) {
	for _, v := range list {
		fmt.Println(v)
	}
}

func TestBst_Add(t *testing.T) {
	/////////////////
	//       5     //
	//     /   \   //
	//    3     8  //
	//   / \   /   //
	//  2   4 6    //
	// /       \   //
	// 1        7  //
	/////////////////

	// 插入数据
	list := []int{5, 8, 3, 6, 4, 2, 7, 1}
	node := Bst{}
	for _, v := range list {
		node.Add(v)
	}
	// 前序遍历
	res1 := []int{5, 3, 2, 1, 4, 8, 6, 7}
	// 中序遍历
	res2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// 后序遍历
	res3 := []int{1, 2, 4, 3, 7, 6, 8, 5}

	if !Same(node, func(b Bst) []int {
		var list []int
		b.PreOrder(func(node *Node) {
			list = append(list, node.Num)
		})
		return list
	}, res1) {
		t.Error("前序遍历错误")
	}

	if !Same(node, func(b Bst) []int {
		var list []int
		b.PreOrder2(func(node *Node) {
			list = append(list, node.Num)
		})
		return list
	}, res1) {
		t.Error("前序遍历错误2")
	}

	if !Same(node, func(b Bst) []int {
		var list []int
		b.InOrder(func(node *Node) {
			list = append(list, node.Num)
		})
		return list
	}, res2) {
		t.Error("中序遍历错误")
	}

	if !Same(node, func(b Bst) []int {
		var list []int
		b.InOrder2(func(node *Node) {
			list = append(list, node.Num)
		})
		//Pt(list)
		return list
	}, res2) {
		t.Error("中序遍历错误2")
	}

	if !Same(node, func(b Bst) []int {
		var list []int
		b.PostOrder(func(node *Node) {
			list = append(list, node.Num)
		})
		return list
	}, res3) {
		t.Error("后序遍历错误")
	}

	if !Same(node, func(b Bst) []int {
		var list []int
		b.PostOrder2(func(node *Node) {
			list = append(list, node.Num)
		})
		//Pt(list)
		return list
	}, res3) {
		t.Error("后序遍历错误2")
	}
}
