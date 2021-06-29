package uf_test

import (
	"fmt"

	"github.com/skeletongo/datastructure/uf"
)

type data struct {
	id int // 每个数据的唯一索引，用索引最大值创建并查集
	// ...
}

func (d *data) UniqueId() int {
	return d.id
}

var collections []data

const num = 10

func init() {
	index := -1
	for i := 0; i < num; i++ {
		index++
		collections = append(collections, data{
			id: index,
		})
	}
}

func ExampleNew() {
	u := uf.New(num)
	fmt.Println(u.IsConnected(collections[0].UniqueId(), collections[num-1].UniqueId()))

	u.Union(collections[0].UniqueId(), collections[num-1].UniqueId())
	fmt.Println(u.IsConnected(collections[0].UniqueId(), collections[num-1].UniqueId()))

	u.Union(collections[0].UniqueId(), collections[1].UniqueId())
	fmt.Println(u.IsConnected(collections[0].UniqueId(), collections[num-1].UniqueId()))
	// Output:
	// false
	// true
	// true
}
