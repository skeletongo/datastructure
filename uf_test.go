package dataStructure

import (
	"math/rand"
	"testing"
	"time"
)

func TestUnionFind(t *testing.T) {
	num := 1000
	uf := NewUF(num)
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			if i == j {
				continue
			}
			if uf.IsConnected(i, j) {
				t.Error("no connected")
				return
			}
		}
	}

	var test [][]int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 500; i++ {
		p := r.Intn(num)
		q := r.Intn(num)
		test = append(test, []int{p, q})
		uf.UnionElements(p, q)
	}

	for _, v := range test {
		if !uf.IsConnected(v[0], v[1]) {
			t.Error("no connected")
			return
		}
	}
}
