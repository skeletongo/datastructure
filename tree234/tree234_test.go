package tree234

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestTree234_Put(t *testing.T) {
	for i := 0; i < 10000; i++ {
		tree := New(func(a, b interface{}) int {
			return a.(int) - b.(int)
		})

		n := rand.Intn(1000)
		arr := rand.Perm(n)
		for i := 0; i < len(arr); i++ {
			tree.Put(arr[i], nil)
			if !tree.isBalanced() {
				t.Fatal("balance error", arr)
			}
			if !tree.Contains(arr[i]) {
				t.Fatal("contains error", arr)
			}
			if tree.GetSize() != i+1 {
				t.Fatal("size error", arr)
			}
		}
	}
}

func TestTree234_RemoveMin(t *testing.T) {
	for i := 0; i < 10000; i++ {
		tree := New(func(a, b interface{}) int {
			return a.(int) - b.(int)
		})

		n := rand.Intn(1000)
		arr := rand.Perm(n)
		for i := 0; i < len(arr); i++ {
			tree.Put(arr[i], nil)
		}

		for i := 0; i < n; i++ {
			tree.RemoveMin()
			if !tree.isBalanced() {
				t.Fatal("balance error", arr)
			}
			if tree.Contains(i) {
				t.Fatal("contains error", arr)
			}
			if tree.GetSize() != n-1-i {
				t.Fatal("size error", arr)
			}
		}
	}
}

func TestTree234_RemoveMax(t *testing.T) {
	for i := 0; i < 10000; i++ {
		tree := New(func(a, b interface{}) int {
			return a.(int) - b.(int)
		})

		n := rand.Intn(1000)
		arr := rand.Perm(n)
		for i := 0; i < len(arr); i++ {
			tree.Put(arr[i], nil)
		}

		for i := n - 1; i >= 0; i-- {
			tree.RemoveMax()
			if !tree.isBalanced() {
				t.Fatal("balance error", arr)
			}
			if tree.Contains(i) {
				t.Fatal("contains error", arr)
			}
			if tree.GetSize() != i {
				t.Fatal("size error", arr)
			}
		}
	}
}

func TestTree234_Remove(t *testing.T) {
	for i := 0; i < 10000; i++ {
		tree := New(func(a, b interface{}) int {
			return a.(int) - b.(int)
		})

		n := rand.Intn(1000)
		arr := rand.Perm(n)
		arr2 := make([]int, n)
		copy(arr2, arr)
		for i := 0; i < len(arr)/2; i++ {
			tree.Put(arr[i], nil)
		}

		delArr := []int{}

		for i := 0; i < n; i++ {
			j := rand.Intn(len(arr2))
			key := arr2[j]
			arr2 = append(arr2[:j], arr2[j+1:]...)

			delArr = append(delArr, key)

			size := tree.GetSize()
			has := tree.Contains(key)
			tree.Remove(key)
			if !tree.isBalanced() {
				t.Fatal("balance error", arr[:len(arr)/2], delArr)
			}
			if tree.Contains(key) {
				t.Fatal("contains error")
			}
			if (has && tree.GetSize() != size-1) || (!has && tree.GetSize() != size) {
				t.Fatal("size error")
			}
		}
	}
}

// 输出删除最小值的过程,辅助调试代码
func TestTree234_RemoveMin_Image(t *testing.T) {
	tree := New(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})

	arr := []int{}
	for _, v := range arr {
		tree.Put(v, nil)
	}

	if err := tree.Img(fmt.Sprintf("%d_removemin", 0)); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(arr); i++ {
		tree.RemoveMin()
		tree.isBalanced()
		if err := tree.Img(fmt.Sprintf("%d_removemin", i+1)); err != nil {
			t.Fatal(err)
		}
	}
}

// 输出删除最大值的过程,辅助调试代码
func TestTree234_RemoveMax_Image(t *testing.T) {
	tree := New(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})

	arr := []int{}
	for _, v := range arr {
		tree.Put(v, nil)
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

// 输出删除节点的过程,辅助调试代码
func TestTree234_Remove_Image(t *testing.T) {
	tree := New(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})

	arr := []int{}
	delArr := []int{}
	for _, v := range arr {
		tree.Put(v, nil)
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
