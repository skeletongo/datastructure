package test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/skeletongo/datastructure/avltree"
	"github.com/skeletongo/datastructure/hashtable"
	"github.com/skeletongo/datastructure/tree23"
	"github.com/skeletongo/datastructure/tree234"
)

/*
avl树，红黑树，哈希表性能测试
*/
func init() {
	rand.Seed(time.Now().UnixNano())
}

func compare(a, b interface{}) int {
	return a.(int) - b.(int)
}

func BenchmarkTree_Contains(b *testing.B) {
	arr := rand.Perm(1000)
	delArr := rand.Perm(1000)
	b.Run("AVLTreeContains", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			tree := avltree.New(compare)
			for i := 0; i < len(arr); i++ {
				tree.Put(arr[i])
			}
			b.StartTimer()
			for i := 0; i < len(delArr); i++ {
				tree.Contains(delArr[i])
			}
		}
	})

	b.Run("Tree23Contains", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			tree := tree23.New(compare)
			for i := 0; i < len(arr); i++ {
				tree.Put(arr[i])
			}
			b.StartTimer()
			for i := 0; i < len(delArr); i++ {
				tree.Contains(delArr[i])
			}
		}
	})

	b.Run("Tree234Contains", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			tree := tree234.New(compare)
			for i := 0; i < len(arr); i++ {
				tree.Put(arr[i])
			}
			b.StartTimer()
			for i := 0; i < len(delArr); i++ {
				tree.Contains(delArr[i])
			}
		}
	})

	b.Run("TreeMapContains", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			tree := hashtable.NewTableTree(compare)
			for i := 0; i < len(arr); i++ {
				tree.Set(arr[i], nil)
			}
			b.StartTimer()
			for i := 0; i < len(delArr); i++ {
				tree.Contains(delArr[i])
			}
		}
	})

	b.Run("GoMapContains", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			m := map[int]struct{}{}
			for i := 0; i < len(arr); i++ {
				m[arr[i]] = struct{}{}
			}
			b.StartTimer()
			for i := 0; i < len(delArr); i++ {
				_, _ = m[delArr[i]]
			}
		}
	})
}

func BenchmarkTree_Put(b *testing.B) {
	arr := rand.Perm(1000)

	b.Run("AVLTreePut", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			tree := avltree.New(compare)
			b.StartTimer()
			for i := 0; i < len(arr); i++ {
				tree.Put(arr[i])
			}
		}
	})

	b.Run("Tree23Put", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			tree := tree23.New(compare)
			b.StartTimer()
			for i := 0; i < len(arr); i++ {
				tree.Put(arr[i])
			}
		}
	})

	b.Run("Tree234Put", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			tree := tree234.New(compare)
			b.StartTimer()
			for i := 0; i < len(arr); i++ {
				tree.Put(arr[i])
			}
		}
	})

	b.Run("TreeMapPut", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			tree := hashtable.NewTableTree(compare)
			b.StartTimer()
			for i := 0; i < len(arr); i++ {
				tree.Set(arr[i], nil)
			}
		}
	})

	b.Run("GoMapPut", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			m := map[int]struct{}{}
			b.StartTimer()
			for i := 0; i < len(arr); i++ {
				m[arr[i]] = struct{}{}
			}
		}
	})
}

func BenchmarkTree_Remove(b *testing.B) {
	arr := rand.Perm(1000)
	delArr := rand.Perm(1000)
	b.Run("AVLTreeRemove", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			tree := avltree.New(compare)
			for i := 0; i < len(arr); i++ {
				tree.Put(arr[i])
			}
			b.StartTimer()
			for i := 0; i < len(delArr); i++ {
				tree.Remove(delArr[i])
			}
		}
	})

	b.Run("Tree23Remove", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			tree := tree23.New(compare)
			for i := 0; i < len(arr); i++ {
				tree.Put(arr[i])
			}
			b.StartTimer()
			for i := 0; i < len(delArr); i++ {
				tree.Remove(delArr[i])
			}
		}
	})

	b.Run("Tree234Remove", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			tree := tree234.New(compare)
			for i := 0; i < len(arr); i++ {
				tree.Put(arr[i])
			}
			b.StartTimer()
			for i := 0; i < len(delArr); i++ {
				tree.Remove(delArr[i])
			}
		}
	})

	b.Run("TreeMapRemove", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			tree := hashtable.NewTableTree(compare)
			for i := 0; i < len(arr); i++ {
				tree.Set(arr[i], nil)
			}
			b.StartTimer()
			for i := 0; i < len(delArr); i++ {
				tree.Remove(delArr[i])
			}
		}
	})

	b.Run("GoMapRemove", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			m := map[int]struct{}{}
			for i := 0; i < len(arr); i++ {
				m[arr[i]] = struct{}{}
			}
			b.StartTimer()
			for i := 0; i < len(delArr); i++ {
				delete(m, delArr[i])
			}
		}
	})
}

func BenchmarkTree_All(b *testing.B) {
	arr := rand.Perm(1000000)
	delArr := rand.Perm(1000000)
	b.Run("AVLTreeAll", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			tree := avltree.New(compare)
			b.StartTimer()
			for i := 0; i < len(arr); i++ {
				tree.Put(arr[i])
			}
			for i := 0; i < len(delArr); i++ {
				tree.Contains(delArr[i])
			}
			for i := 0; i < len(delArr); i++ {
				tree.Remove(delArr[i])
			}
		}
	})

	b.Run("Tree23All", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			tree := tree23.New(compare)
			b.StartTimer()
			for i := 0; i < len(arr); i++ {
				tree.Put(arr[i])
			}
			for i := 0; i < len(delArr); i++ {
				tree.Contains(delArr[i])
			}
			for i := 0; i < len(delArr); i++ {
				tree.Remove(delArr[i])
			}
		}
	})

	b.Run("Tree234All", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			tree := tree234.New(compare)
			b.StartTimer()
			for i := 0; i < len(arr); i++ {
				tree.Put(arr[i])
			}
			for i := 0; i < len(delArr); i++ {
				tree.Contains(delArr[i])
			}
			for i := 0; i < len(delArr); i++ {
				tree.Remove(delArr[i])
			}
		}
	})

	b.Run("TreeMapAll", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			tree := hashtable.NewTableTree(compare)
			b.StartTimer()
			for i := 0; i < len(arr); i++ {
				tree.Set(arr[i], nil)
			}
			for i := 0; i < len(delArr); i++ {
				tree.Contains(delArr[i])
			}
			for i := 0; i < len(delArr); i++ {
				tree.Remove(delArr[i])
			}
		}
	})

	b.Run("GoMapAll", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			m := map[int]struct{}{}
			b.StartTimer()
			for i := 0; i < len(arr); i++ {
				m[arr[i]] = struct{}{}
			}
			for i := 0; i < len(delArr); i++ {
				_, _ = m[delArr[i]]
			}
			for i := 0; i < len(delArr); i++ {
				delete(m, delArr[i])
			}
		}
	})
}
