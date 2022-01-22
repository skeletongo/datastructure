package btree_test

import (
	"fmt"

	"github.com/skeletongo/datastructure/btree"
)

type kv struct {
	key   int
	value int
}

func ExampleBTree() {
	tree := btree.New(3, func(a, b interface{}) int {
		return a.(*kv).key - b.(*kv).key
	})
	// add
	tree.Put(&kv{key: 1, value: -1})
	fmt.Println(tree)
	tree.Put(&kv{key: 2, value: -2})
	fmt.Println(tree)
	tree.Put(&kv{key: 3, value: -3})
	fmt.Println(tree)
	tree.Put(&kv{key: 4, value: -4})
	fmt.Println(tree)
	tree.Put(&kv{key: 5, value: -5})
	fmt.Println(tree)
	tree.Put(&kv{key: 6, value: -6})
	fmt.Println(tree)
	tree.Put(&kv{key: 7, value: -7})
	fmt.Println(tree)

	// remove
	tree.Remove(&kv{key: 2})
	tree.Remove(&kv{key: 4})
	tree.Remove(&kv{key: 6})
	fmt.Println(tree)

	// modify
	tree.Put(&kv{key: 1, value: 1})
	tree.Put(&kv{key: 3, value: 3})
	tree.Put(&kv{key: 5, value: 5})
	tree.Put(&kv{key: 7, value: 7})
	fmt.Println(tree)

	// find
	fmt.Println(tree.Get(&kv{key: 1}))
	fmt.Println(tree.Get(&kv{key: 3}))
	fmt.Println(tree.Get(&kv{key: 5}))
	fmt.Println(tree.Get(&kv{key: 7}))
	fmt.Println(tree.Get(&kv{key: 0}))

	// Output:
	// &{1 -1}
	//
	// &{1 -1},&{2 -2}
	//
	// &{2 -2}
	// |--&{1 -1}
	// `--&{3 -3}
	//
	// &{2 -2}
	// |--&{1 -1}
	// `--&{3 -3},&{4 -4}
	//
	// &{2 -2},&{4 -4}
	// |--&{1 -1}
	// |--&{3 -3}
	// `--&{5 -5}
	//
	// &{2 -2},&{4 -4}
	// |--&{1 -1}
	// |--&{3 -3}
	// `--&{5 -5},&{6 -6}
	//
	// &{4 -4}
	// |--&{2 -2}
	// |  |--&{1 -1}
	// |  `--&{3 -3}
	// `--&{6 -6}
	//    |--&{5 -5}
	//    `--&{7 -7}
	//
	// &{5 -5}
	// |--&{1 -1},&{3 -3}
	// `--&{7 -7}
	//
	// &{5 5}
	// |--&{1 1},&{3 3}
	// `--&{7 7}
	//
	// &{1 1}
	// &{3 3}
	// &{5 5}
	// &{7 7}
	// <nil>
}
