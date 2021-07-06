package bst

import (
	"fmt"
	"github.com/skeletongo/datastructure/common"
)

type node struct {
	left, right *node
	key, value  interface{}
}

func newNode(key, value interface{}) *node {
	return &node{key: key, value: value}
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
