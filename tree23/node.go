package tree23

import (
	"github.com/skeletongo/datastructure/common"
)

const (
	Red   = true
	Black = false
)

type node struct {
	left, right *node
	color       bool
	n           int         // 以当前节点为根的树的节点数量
	key, value  interface{} // 存储映射的键值对
}

func newNode(key, value interface{}) *node {
	return &node{
		key:   key,
		value: value,
		color: Red,
		n:     1,
	}
}

func (n *node) GetLeftNode() common.INodeKey {
	return n.left
}

func (n *node) GetRightNode() common.INodeKey {
	return n.right
}

func (n *node) GetKey() interface{} {
	return n.key
}

func (n *node) GetValue() interface{} {
	return n.value
}
