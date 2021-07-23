package tree234

import (
	"fmt"

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
	value       interface{} // 存储映射的键值对
}

func newNode(value interface{}) *node {
	return &node{
		value: value,
		color: Red,
		n:     1,
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
	if isRed(n) {
		return ",color=red,style=filled,fontcolor=white"
	}
	return ",color=black,style=filled,fontcolor=white"
}

func (n *node) String() string {
	return fmt.Sprintf("%v", n.value)
}

// isRed 是否为红节点
func isRed(n *node) bool {
	if n == nil {
		return Black
	}
	return n.color // n.color == red 不用等号判断需要定义Red为true
}
