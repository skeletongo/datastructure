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
	common.TreeNode
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

func (n *node) GetColor() string {
	if isRed(n) {
		return "red"
	}
	return "black"
}

func (n *node) GetLeftTreeNode() common.ITreeNode {
	return n.left
}

func (n *node) GetRightTreeNode() common.ITreeNode {
	return n.right
}

// BuildIndex 以层序遍历的顺序给节点设置递增编号，这个是画svg图用的
func (n *node) BuildIndex() {
	i := 0
	common.LevelOrder(n, func(node common.INode) {
		treeNode := node.(common.ITreeNode)
		treeNode.SetIndex(i)
		i++
	})
}

func (n *node) String() string {
	return fmt.Sprintf("%v: %v", n.key, n.value)
}

// isRed 是否为红节点
func isRed(n *node) bool {
	if n == nil {
		return Black
	}
	return n.color // n.color == red 不用等号判断需要定义Red为true
}
