package avltree

import (
	"fmt"
	"github.com/skeletongo/datastructure/common"
)

type node struct {
	height      int // 以当前节点为根的树的高度
	left, right *node
	key, value  interface{} // 存储映射的键值对
}

func newNode(key, value interface{}) *node {
	return &node{
		height: 1,
		key:    key,
		value:  value,
	}
}

func (n *node) GetLeftNode() common.INode {
	return n.left
}

func (n *node) GetRightNode() common.INode {
	return n.right
}

func (n *node) GetKey() interface{} {
	return n.key
}

func (n *node) GetValue() interface{} {
	return n.value
}

func (n *node) String() string {
	return fmt.Sprintf("%v: %v", n.key, n.value)
}
