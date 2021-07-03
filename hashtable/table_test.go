package hashtable

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestTableList(t *testing.T) {
	n := 1000
	arr := rand.Perm(n)
	tb := NewTableList()

	if !tb.IsEmpty() {
		t.Error("IsEmpty() != true error")
	}

	var m = make(map[int]int)
	for k, v := range arr {
		m[v] = k
		tb.Set(v, k)
	}

	testFunc := func() {
		ids := rand.Perm(2 * len(m))
		for _, v := range ids {
			_, ok := m[v]
			if tb.Contains(v) != ok {
				t.Error("Contains() error")
			}
			if (tb.Get(v) != nil) != ok {
				t.Error("Get() error")
			}
		}

		if tb.IsEmpty() != (len(m) == 0) {
			t.Error("IsEmpty() error")
		}
		if tb.GetSize() != len(m) {
			t.Error("GetSize() error")
		}
		for k, v := range m {
			if tb.Get(k) != v {
				t.Error("Get(v) error")
			}
		}
	}

	testFunc()

	// 修改
	for k, v := range arr {
		tb.Set(k, v)
		m[k] = v
	}
	testFunc()

	// 删除
	for i := 0; i < n; i++ {
		if rand.Intn(10) < 5 {
			for k := range m {
				tb.Remove(k)
				delete(m, k)
				break
			}
			if rand.Intn(10) < 5 {
				testFunc()
			}
		}
	}
	testFunc()
}

func TestTableTree(t *testing.T) {
	n := 1000
	arr := rand.Perm(n)
	tb := NewTableTree(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})

	if !tb.IsEmpty() {
		t.Error("IsEmpty() != true error")
	}

	var m = make(map[int]int)
	for k, v := range arr {
		m[v] = k
		tb.Set(v, k)
	}

	testFunc := func() {
		ids := rand.Perm(2 * len(m))
		for _, v := range ids {
			_, ok := m[v]
			if tb.Contains(v) != ok {
				t.Error("Contains() error")
			}
			if (tb.Get(v) != nil) != ok {
				t.Error("Get() error")
			}
		}

		if tb.IsEmpty() != (len(m) == 0) {
			t.Error("IsEmpty() error")
		}
		if tb.GetSize() != len(m) {
			t.Error("GetSize() error")
		}
		for k, v := range m {
			if tb.Get(k) != v {
				t.Error("Get(v) error")
			}
		}
	}

	testFunc()

	// 修改
	for k, v := range arr {
		tb.Set(k, v)
		m[k] = v
	}
	testFunc()

	// 删除
	for len(m) > 0 {
		for k := range m {
			tb.Remove(k)
			delete(m, k)
			testFunc()
		}
	}
}
