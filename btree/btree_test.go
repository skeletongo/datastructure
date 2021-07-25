package btree

import (
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
			if tree.GetSize() != i+1 {
				t.Fatal("size error", arr)
			}
		}
	}
}
