package dataStructure

import (
	"fmt"
	"testing"
)

/////////////////
//       5     //
//     /   \   //
//    3     8  //
//   / \   /   //
//  2   4 6    //
// /       \   //
// 1        7  //
/////////////////

// 二分搜索树
var node Bst

// 插入数据
var list = []int{5, 8, 3, 6, 4, 2, 7, 1}

// 前序遍历
var res1 = []int{5, 3, 2, 1, 4, 8, 6, 7}

// 中序遍历
var res2 = []int{1, 2, 3, 4, 5, 6, 7, 8}

// 后序遍历
var res3 = []int{1, 2, 4, 3, 7, 6, 8, 5}

// 层序遍历
var res4 = []int{5, 3, 8, 2, 4, 6, 1, 7}

func CreateBst() Bst {
	node = Bst{}
	for _, v := range list {
		node.Add2(v)
	}
	return node
}

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
		fmt.Print(v, " ")
	}
}

func TestBst(t *testing.T) {
	node = CreateBst()
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
		//Pt(list)
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
		//Pt(list)
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

	if !Same(node, func(bst Bst) []int {
		var list []int
		bst.LevelOrder(func(node *Node) {
			list = append(list, node.Num)
		})
		//Pt(list)
		return list
	}, res4) {
		t.Error("层序遍历错误")
	}
}

func TestBst_DeleteMin(t *testing.T) {
	node = CreateBst()
	for _, v := range res2 {
		if node.RemoveMin() != v {
			t.Error("删除最小数错误")
		}
	}
}

func TestBst_DeleteMax(t *testing.T) {
	node = CreateBst()
	for i := len(res2) - 1; i >= 0; i-- {
		if node.RemoveMax() != res2[i] {
			t.Error("删除最大数错误")
		}
	}
}

func TestBst_Remove(t *testing.T) {
	node = CreateBst()
	data := []int{0, 5, 1, 3, 7, 4, 6, 8, 2}
	res := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 6, 7, 8},
		{2, 3, 4, 6, 7, 8},
		{2, 4, 6, 7, 8},
		{2, 4, 6, 8},
		{2, 6, 8},
		{2, 8},
		{2},
		{},
	}
	for k, v := range data {
		node.Remove(v)
		if !Same(node, func(bst Bst) []int {
			var list []int
			bst.InOrder(func(i *Node) {
				list = append(list, i.Num)
			})
			//Pt(list)
			return list
		}, res[k]) {
			t.Error("删除任一元素错误")
		}
		//fmt.Println()
	}
}
