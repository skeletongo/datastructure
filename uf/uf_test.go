package uf

import (
	"errors"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// initUF 用已知集合创建并查集
func initUF(u *UF, union []map[int]struct{}) {
	for _, v := range union {
		var a = -1
		for k := range v {
			if a < 0 {
				a = k
				continue
			}
			u.Union(a, k)
		}
	}
}

// newUF 随机创建并查集
// 返回结果切片中每一个map中的数据算是同一个集合中的数据，切片元素数量就是有多少个不相连的集合
func newUF(u *UF) []map[int]struct{} {
	n := u.GetSize()
	data := rand.Perm(n)
	var union []map[int]struct{}
	var col map[int]struct{}
	for len(data) > 0 {
		if rand.Intn(4) == 0 {
			col = make(map[int]struct{})
			union = append(union, col)
		} else {
			if len(union) == 0 {
				continue
			}
			col = union[len(union)-1]
		}
		col[data[0]] = struct{}{}
		data = data[1:]
	}
	initUF(u, union)
	return union
}

// testUF 测试并查集是否正常
func testUF(u *UF) error {
	n := u.GetSize()
	// 随机创建并查集
	union := newUF(u)
	//fmt.Printf("--> union map: %v \n",union)
	// 验证
	var col map[int]struct{}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			for _, v := range union {
				if _, ok := v[i]; ok {
					col = v
					break
				}
			}
			_, ok := col[j]
			//fmt.Println(i, j, ok)
			if u.IsConnected(i, j) != ok {
				return errors.New("IsConnected error")
			}
		}
	}
	return nil
}

func TestFindParentFuncLess(t *testing.T) {
	n := 1000
	u := NewFunc(n, FindParentFuncLess)
	if err := testUF(u); err != nil {
		t.Error(err)
	}
}

func TestFindParentFuncMore(t *testing.T) {
	n := 1000
	u := NewFunc(n, FindParentFuncMore)
	if err := testUF(u); err != nil {
		t.Error(err)
	}
}

func BenchmarkUF_IsConnected(b *testing.B) {
	n := 1000
	uf := New(n)
	newUF(uf)

	f := func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for i := 0; i < n; i++ {
				for j := i; j < n; j++ {
					uf.IsConnected(i, j)
				}
			}
		}
	}

	uf.Set(FindParentFuncLess)
	b.Run("FindParentFuncLess", f)

	uf.Set(FindParentFuncMore)
	b.Run("FindParentFuncMore", f)

	//goos: windows
	//goarch: amd64
	//pkg: dataStructure/uf
	//cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
	//BenchmarkUF_IsConnected
	//BenchmarkUF_IsConnected/FindParentFuncLess
	//BenchmarkUF_IsConnected/FindParentFuncLess-8         	     213	   5627625 ns/op
	//BenchmarkUF_IsConnected/FindParentFuncMore
	//BenchmarkUF_IsConnected/FindParentFuncMore-8         	     170	   6955569 ns/op
}

func BenchmarkUF_Union(b *testing.B) {
	n := 1000
	arr := rand.Perm(n)
	findFunc := FindParentFuncLess

	f := func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			uf := NewFunc(n, findFunc)
			b.StartTimer()
			for j := 0; j < n; j += 2 {
				uf.Union(arr[j], arr[j+1])
			}
		}
	}

	b.Run("FindParentFuncLess", f)

	findFunc = FindParentFuncMore
	b.Run("FindParentFuncMore", f)

	//goos: windows
	//goarch: amd64
	//pkg: dataStructure/uf
	//cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
	//BenchmarkUF_Union
	//BenchmarkUF_Union/FindParentFuncLess
	//BenchmarkUF_Union/FindParentFuncLess-8         	  234108	      4940 ns/op
	//BenchmarkUF_Union/FindParentFuncMore
	//BenchmarkUF_Union/FindParentFuncMore-8         	  213864	      5192 ns/op
}
