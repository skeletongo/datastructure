package bst

import (
	"fmt"

	"github.com/skeletongo/datastructure/common"
)

type node struct {
	left, right *node
	value       interface{}
}

func newNode(value interface{}) *node {
	return &node{value: value}
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
