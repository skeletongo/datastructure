package bst

import (
	"testing"
)

func TestBST(t *testing.T) {
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
	var list = []int{5, 8, 3, 6, 4, 2, 7, 1}

	// 前序遍历
	var res1 = []int{5, 3, 2, 1, 4, 8, 6, 7}

	// 中序遍历
	var res2 = []int{1, 2, 3, 4, 5, 6, 7, 8}

	// 后序遍历
	var res3 = []int{1, 2, 4, 3, 7, 6, 8, 5}

	// 层序遍历
	var res4 = []int{5, 3, 8, 2, 4, 6, 1, 7}

	var res5 = []int{2, 4, 5, 7}

	bst := NewBST(func(a, b interface{}) int {
		m := a.(int)
		n := b.(int)
		return m - n
	})

	for _, v := range list {
		bst.Add(v)
	}

	n := 0
	bst.PreOrder(func(e interface{}) {
		if res1[n] != e {
			t.Error("前序遍历错误")
		}
		n++
	})

	n = 0
	bst.InOrder(func(e interface{}) {
		if res2[n] != e {
			t.Error("中序遍历错误")
		}
		n++
	})

	n = 0
	bst.PostOrder(func(e interface{}) {
		if res3[n] != e {
			t.Error("后序遍历错误")
		}
		n++
	})

	n = 0
	bst.DFS(func(e interface{}) {
		if res2[n] != e {
			t.Error("深度优先遍历错误", e)
		}
		n++
	})

	n = 0
	bst.BFS(func(e interface{}) {
		if res4[n] != e {
			t.Error("层序遍历错误")
		}
		n++
	})

	if n := bst.RemoveMin(); n != 1 {
		t.Error("删除最小值错误", n)
	}

	if n := bst.RemoveMax(); n != 8 {
		t.Error("删除最大值错误", n)
	}

	if n := bst.GetSize(); n != 6 {
		t.Error("节点个数统计错误", n)
	}

	if !bst.Remove(3) {
		t.Error("删除任意元素错误", 3)
	}

	if bst.Contains(3) || !bst.Contains(4) {
		t.Error("查询包含指定元素错误")
	}

	if !bst.Remove(6) {
		t.Error("删除任意元素错误", 6)
	}
	if bst.Remove(8) {
		t.Error("删除任意元素错误", 8)
	}
	n = 0
	bst.InOrder(func(e interface{}) {
		if res5[n] != e {
			t.Error("删除任意元素错误", -1)
		}
		n++
	})
}
