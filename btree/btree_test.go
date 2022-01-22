package btree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestBTree_Put(t *testing.T) {
	for i := 0; i < 10000; i++ {
		tree := New(3+rand.Intn(3), func(a, b interface{}) int {
			return a.(int) - b.(int)
		})

		n := rand.Intn(1000)
		arr := rand.Perm(n)
		for i := 0; i < len(arr); i++ {
			tree.Put(arr[i])
			if !tree.isBtree() {
				t.Fatal("isBtree error", arr, tree.Rank())
			}
			if !tree.Contains(arr[i]) {
				t.Fatal("contains error", arr)
			}
		}
	}
}

func TestBTree_RemoveMin(t *testing.T) {
	for i := 0; i < 10000; i++ {
		tree := New(3+rand.Intn(3), func(a, b interface{}) int {
			return a.(int) - b.(int)
		})

		n := rand.Intn(1000)
		arr := rand.Perm(n)
		for i := 0; i < len(arr); i++ {
			tree.Put(arr[i])
		}

		for i := 0; i < n; i++ {
			tree.RemoveMin()
			if !tree.isBtree() {
				t.Fatalf("isBtree error arr: %v d: %d", arr, tree.d)
			}
			if tree.Contains(i) {
				t.Fatalf("contains error arr: %v d: %d", arr, tree.d)
			}
		}
	}
}

// 输出删除最小值的过程,辅助调试代码
func TestBTree_RemoveMin_Image(t *testing.T) {
	tree := New(5, func(a, b interface{}) int {
		return a.(int) - b.(int)
	})

	arr := []int{}
	for _, v := range arr {
		tree.Put(v)
	}

	if err := tree.Img(fmt.Sprintf("%d_removemin", 0)); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(arr); i++ {
		tree.RemoveMin()
		if err := tree.Img(fmt.Sprintf("%d_removemin", i+1)); err != nil {
			t.Fatal(err)
		}
	}
}

func TestBTree_RemoveMax(t *testing.T) {
	for i := 0; i < 10000; i++ {
		tree := New(3+rand.Intn(3), func(a, b interface{}) int {
			return a.(int) - b.(int)
		})

		n := rand.Intn(1000)
		arr := rand.Perm(n)
		for i := 0; i < len(arr); i++ {
			tree.Put(arr[i])
		}

		for i := n - 1; i >= 0; i-- {
			tree.RemoveMax()
			if !tree.isBtree() {
				t.Fatalf("isBtree error arr: %v d: %d", arr, tree.d)
			}
			if tree.Contains(i) {
				t.Fatalf("contains error arr: %v d: %d", arr, tree.d)
			}
		}
	}
}

// 输出删除最大值的过程,辅助调试代码
func TestBTree_RemoveMax_Image(t *testing.T) {
	tree := New(3, func(a, b interface{}) int {
		return a.(int) - b.(int)
	})

	arr := []int{3, 5, 1, 4, 0, 2, 6}
	for _, v := range arr {
		tree.Put(v)
	}

	if err := tree.Img(fmt.Sprintf("%d_removemax", 0)); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(arr); i++ {
		tree.RemoveMax()
		if err := tree.Img(fmt.Sprintf("%d_removemax", i+1)); err != nil {
			t.Fatal(err)
		}
	}
}

func TestBTree_Remove(t *testing.T) {
	for i := 0; i < 10000; i++ {
		tree := New(3, func(a, b interface{}) int {
			return a.(int) - b.(int)
		})

		n := rand.Intn(1000)
		arr := rand.Perm(n)
		arr2 := make([]int, n)
		copy(arr2, arr)
		for i := 0; i < len(arr)/2; i++ {
			tree.Put(arr[i])
		}

		delArr := []int{}

		for i := 0; i < n; i++ {
			j := rand.Intn(len(arr2))
			key := arr2[j]
			arr2 = append(arr2[:j], arr2[j+1:]...)

			delArr = append(delArr, key)

			tree.Remove(key)
			if !tree.isBtree() {
				t.Fatalf("isBtree error arr:%v delArr:%v d:%v", arr[:len(arr)/2], delArr, tree.d)
			}
			if tree.Contains(key) {
				t.Fatal("contains error")
			}
		}
	}
}

// 输出删除节点的过程,辅助调试代码
func TestBTree_Remove_Image(t *testing.T) {
	tree := New(3, func(a, b interface{}) int {
		return a.(int) - b.(int)
	})

	arr := []int{}
	delArr := []int{}
	for _, v := range arr {
		tree.Put(v)
	}

	if err := tree.Img(fmt.Sprintf("%d_remove", 0)); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(delArr); i++ {
		tree.Remove(delArr[i])
		if err := tree.Img(fmt.Sprintf("%d_remove_%d", i+1, delArr[i])); err != nil {
			t.Fatal(err)
		}
	}
}
