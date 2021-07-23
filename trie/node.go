package trie

import (
	"fmt"
	"github.com/skeletongo/datastructure/common"
)

type node struct {
	isWord bool
	next   map[rune]*node // 使用rune类型可以支持多种语言而非只支持英文
	value  interface{}
}

func (n *node) GetChildren() []common.TreeNode {
	var arr []common.TreeNode
	for _, v := range n.next {
		arr = append(arr, v)
	}
	return arr
}

func (n *node) GetValue() interface{} {
	return n.value
}

func (n *node) String() string {
	return fmt.Sprint(n.value)
}

func newNode() *node {
	return &node{
		isWord: false,
		next:   make(map[rune]*node),
	}
}
