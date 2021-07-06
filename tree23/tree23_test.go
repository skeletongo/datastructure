package tree23

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestTree23_RemoveMin(t *testing.T) {
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
			if !tree.Contains(i) && tree.GetSize() != n-i {
				t.Fatal("Put error")
			}
			tree.RemoveMin()
			tree.isBalanced()
			if tree.Contains(i) || tree.GetSize() != n-i-1 {
				t.Fatal("RemoveMin error")
			}
		}
	}
}

func TestTree23_RemoveMax(t *testing.T) {
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
			if !tree.Contains(i) && tree.GetSize() != i+1 {
				t.Fatal("Put error")
			}
			tree.RemoveMax()
			tree.isBalanced()
			if tree.Contains(i) || tree.GetSize() != i {
				t.Fatal("RemoveMax error")
			}
		}
	}
}

func TestTree23_Remove(t *testing.T) {
	for i := 0; i < 10000; i++ {
		fmt.Println("------------------")
		tree := New(func(a, b interface{}) int {
			return a.(int) - b.(int)
		})
		n := rand.Intn(9)
		arr := rand.Perm(n)
		for i := 0; i < len(arr); i++ {
			tree.Put(arr[i], nil)
		}

		for i := 0; i < 2*n; i++ {
			key := rand.Intn(n)
			if rand.Intn(2) == 0 {
				key = -key
			}
			has := tree.Contains(key)
			size := tree.GetSize()
			fmt.Println(key)
			fmt.Println(tree)
			tree.Remove(key)
			if has {
				if tree.Contains(key) {
					t.Fatal("Remove error")
				}
				if tree.GetSize()+1 != size {
					//t.Fatal("Remove GetSize error", tree.GetSize(), size)
				}
			} else {
				if tree.Contains(key) {
					t.Fatal("Contains error")
				}
				if tree.GetSize() != size {
					//t.Fatal("Remove GetSize error", tree.GetSize(), size)
				}
			}
		}
	}
}

func TestTree23_String(t *testing.T) {
	tree := New(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	arr := rand.Perm(10)
	for i := 0; i < len(arr); i++ {
		tree.Put(arr[i], nil)
	}
	fmt.Println(tree)
}

func TestPrintSvg(t *testing.T) {
	tree := New(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	arr := rand.Perm(10)
	for i := 0; i < len(arr); i++ {
		tree.Put(arr[i], nil)
	}
	tree.Svg("")
}
