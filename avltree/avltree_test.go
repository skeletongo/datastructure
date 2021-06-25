package avltree

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestAVLTree(t *testing.T) {
	n := 1000
	arr := rand.Perm(n)
	tree := New(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})

	if !tree.IsEmpty() {
		t.Error("IsEmpty() != true error")
	}

	var m = make(map[int]int)
	for k, v := range arr {
		m[v] = k
		tree.Put(v, k)
	}

	testFunc := func() {
		ids := rand.Perm(2 * len(m))
		for _, v := range ids {
			_, ok := m[v]
			if tree.Contains(v) != ok {
				t.Error("Contains() error")
			}
			if (tree.Get(v) != nil) != ok {
				t.Error("Get() error")
			}
		}

		if tree.IsEmpty() != (len(m) == 0) {
			t.Error("IsEmpty() error")
		}
		if tree.GetSize() != len(m) {
			t.Error("GetSize() error")
		}
		if !tree.isBST() {
			t.Error("isBST() error")
		}
		if !tree.isBalanced() {
			t.Error("isBalanced() error")
		}
		for k, v := range m {
			if tree.Get(k) != v {
				t.Error("Get(v) error")
			}
		}
	}

	testFunc()

	for k, v := range arr {
		tree.Put(k, v)
		m[k] = v
	}
	testFunc()

	for i := 0; i < n; i++ {
		if rand.Intn(10) < 5 {
			for k := range m {
				tree.Remove(k)
				delete(m, k)
				break
			}
			if rand.Intn(10) < 5 {
				testFunc()
			}
		}
	}
	testFunc()

	for k, v := range arr {
		tree.Put(k, v)
		m[k] = v
	}
	testFunc()
}
