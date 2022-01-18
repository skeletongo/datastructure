package btree

import (
	"fmt"
	"strings"

	"github.com/skeletongo/datastructure/common"
)

type node struct {
	n        int           // 以当前节点为根的子树中的键值对数量
	parent   *node         // 父节点
	values   []interface{} // 所有元素
	children []*node       // 所有子节点
}

// count 计算元素数量
func (n *node) count() {
	n.n = len(n.values)
	for _, v := range n.children {
		n.n += v.n
	}
}

func (n *node) GetChildren() []common.TreeNode {
	var arr []common.TreeNode
	for _, v := range n.children {
		arr = append(arr, v)
	}
	return arr
}

func (n *node) GetValue() interface{} {
	return n.values
}

func (n *node) Children() []common.BTreeNodeGraph {
	var arr []common.BTreeNodeGraph
	for _, v := range n.children {
		arr = append(arr, v)
	}
	return arr
}

func (n *node) Values() []interface{} {
	return n.values
}

func (n *node) String() string {
	var arr []string
	for _, v := range n.values {
		arr = append(arr, fmt.Sprint(v))
	}
	return strings.Join(arr, ",")
}
