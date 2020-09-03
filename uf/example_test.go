package uf_test

import (
	"dataStructure/uf"
	"fmt"
)

type data struct {
	id int
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
	fmt.Println(u.IsConnected(&collections[0], &collections[num-1]))
	u.UnionElements(&collections[0], &collections[num-1])
	fmt.Println(u.IsConnected(&collections[0], &collections[num-1]))
	u.UnionElements(&collections[0], &collections[1])
	fmt.Println(u.IsConnected(&collections[0], &collections[num-1]))
	// Output:
	// false
	// true
	// true
}
