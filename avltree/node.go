package avltree

import (
	"fmt"

	"github.com/skeletongo/datastructure/common"
)

type node struct {
	height      int // 以当前节点为根的树的高度
	left, right *node
	value       interface{} // 存储映射的键值对
}

func newNode(value interface{}) *node {
	return &node{
		height: 1,
		value:  value,
	}
}

func (n *node) GetLeftNode() common.BSTNode {
	return n.left
}

func (n *node) GetRightNode() common.BSTNode {
	return n.right
}

func (n *node) GetValue() interface{} {
	return n.value
}

func (n *node) LeftNode() common.BSTNodeGraph {
	return n.left
}

func (n *node) RightNode() common.BSTNodeGraph {
	return n.right
}

func (n *node) GetAttribute() string {
	return ""
}

func (n *node) String() string {
	return fmt.Sprintf("%v", n.value)
}
