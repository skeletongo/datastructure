package bst

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
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
var list = []int{5, 8, 3, 6, 4, 2, 7, 1}

var testBST = New(func(a, b interface{}) int {
	m := a.(int)
	n := b.(int)
	return m - n
})

func init() {
	rand.Seed(time.Now().UnixNano())
	for _, v := range list {
		testBST.Add(v)
	}
}

func TestBST(t *testing.T) {
	// 前序遍历
	var res1 = []int{5, 3, 2, 1, 4, 8, 6, 7}

	// 中序遍历
	var res2 = []int{1, 2, 3, 4, 5, 6, 7, 8}

	// 后序遍历
	var res3 = []int{1, 2, 4, 3, 7, 6, 8, 5}

	// 层序遍历
	var res4 = []int{5, 3, 8, 2, 4, 6, 1, 7}

	bst := New(func(a, b interface{}) int {
		m := a.(int)
		n := b.(int)
		return m - n
	})

	for _, v := range list {
		bst.Add(v)
	}

	n := 0
	bst.PreOrder(func(v interface{}) {
		//fmt.Println(v)
		if res1[n] != v.(int) {
			t.Error("PreOrder error")
		}
		n++
	})

	n = 0
	bst.InOrder(func(v interface{}) {
		if res2[n] != v.(int) {
			t.Error("InOrder error")
		}
		n++
	})

	n = 0
	bst.PostOrder(func(v interface{}) {
		if res3[n] != v.(int) {
			t.Error("PostOrder error")
		}
		n++
	})

	n = 0
	for _, v := range list {
		if !bst.Contains(v) {
			t.Error("Contains error")
		}
		if bst.Contains(0) {
			t.Error("Contains error")
		}
	}

	bst = New(func(a, b interface{}) int {
		m := a.(int)
		n := b.(int)
		return m - n
	})

	for _, v := range list {
		bst.AddNR(v)
	}

	n = 0
	bst.PreOrderNR(func(v interface{}) {
		//fmt.Println(v)
		if res1[n] != v.(int) {
			t.Error("PreOrderNR error")
		}
		n++
	})

	n = 0
	bst.InOrderNR(func(v interface{}) {
		if res2[n] != v.(int) {
			t.Error("InOrderNR error")
		}
		n++
	})

	n = 0
	bst.PostOrderNR(func(v interface{}) {
		if res3[n] != v.(int) {
			t.Error("PostOrderNR error")
		}
		n++
	})

	n = 0
	bst.PreOrderNRC(func(v interface{}) {
		//fmt.Println(v)
		if res1[n] != v.(int) {
			t.Error("PreOrderNRC error")
		}
		n++
	})

	n = 0
	bst.InOrderNRC(func(v interface{}) {
		if res2[n] != v.(int) {
			t.Error("InOrderNRC error")
		}
		n++
	})

	n = 0
	bst.PostOrderNRC(func(v interface{}) {
		if res3[n] != v.(int) {
			t.Error("PostOrderNRC error")
		}
		n++
	})

	n = 0
	for _, v := range list {
		if !bst.ContainsNR(v) {
			t.Error("ContainsNR error")
		}
		if bst.ContainsNR(0) {
			t.Error("ContainsNR error")
		}
	}

	// LevelOrder
	n = 0
	bst.LevelOrder(func(v interface{}) {
		if res4[n] != v.(int) {
			t.Error("LevelOrder error")
		}
		n++
	})

	// RemoveMin RemoveMax

	if v := bst.RemoveMin(); v != 1 {
		t.Error("RemoveMin", v)
	}

	if v := bst.RemoveMax(); v != 8 {
		t.Error("RemoveMax", v)
	}

	if n = bst.GetSize(); n != 6 {
		t.Error("GetSize error", n)
	}

	bst.Add(1)
	bst.Add(8)

	if v := bst.RemoveMinNR(); v != 1 {
		t.Error("RemoveMinNR", v)
	}

	if v := bst.RemoveMaxNR(); v != 8 {
		t.Error("RemoveMaxNR", v)
	}

	if n = bst.GetSize(); n != 6 {
		t.Error("GetSize error", n)
	}

	bst.Add(1)
	bst.Add(8)

	// Remove

	bst.Remove(1)
	if n = bst.GetSize(); n != 7 {
		t.Error("Remove error", n)
	}
	if bst.Contains(1) || !bst.Contains(2) {
		t.Error("Remove error")
	}
	bst.Add(1)

	bst.Remove(2)
	if n = bst.GetSize(); n != 7 {
		t.Error("Remove error", n)
	}
	if bst.Contains(2) || !bst.Contains(1) {
		t.Error("Remove error")
	}
	bst.Add(2)

	bst.Remove(6)
	if n = bst.GetSize(); n != 7 {
		t.Error("Remove error", n)
	}
	if bst.Contains(6) || !bst.Contains(7) {
		t.Error("Remove error")
	}
	bst.Add(6)

	bst.Remove(3)
	if n = bst.GetSize(); n != 7 {
		t.Error("Remove error", n)
	}
	if bst.Contains(3) || !bst.Contains(4) || !bst.Contains(4) {
		t.Error("Remove error")
	}
	bst.Add(3)

	for _, v := range list {
		bst.Set(v)
	}
	for _, v := range list {
		if bst.Get(v) != v {
			t.Error("Set/Get error \n", bst.String())
		}
	}
	if bst.Get(100) != nil {
		t.Error("Set/Get error \n", bst.String())
	}
}

func TestPrintImg(t *testing.T) {
	bst := New(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})

	arr := rand.Perm(20)
	for i := 0; i < len(arr); i++ {
		bst.Add(arr[i])
	}
	bst.Img("")
}

func TestNode_String(t *testing.T) {
	bst := New(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})

	arr := rand.Perm(10)
	for i := 0; i < len(arr); i++ {
		bst.Add(arr[i])
	}
	fmt.Println(bst)
}
