package common

import (
	"bytes"
	"fmt"
)

// TreeNode 多叉树
type TreeNode interface {
	GetChildren() []TreeNode
	GetValue() interface{}
}

// PrePrintTree 前序打印多叉树
func PrePrintTree(n TreeNode) string {
	if IsNil(n) {
		return ""
	}
	buf := &bytes.Buffer{}
	prePrintTree(buf, "", n, 0, false)
	return buf.String()
}

func prePrintTree(buf *bytes.Buffer, prefix string, node TreeNode, n int, end bool) {
	// 接口值包含两部分，所包含的动态类型和动态类型的值
	// 只有两部分都为nil，接口值才等于nil
	if IsNil(node) {
		return
	}
	if n > 0 {
		if end {
			prefix += "`--"
		} else {
			prefix += "|--"
		}
	}
	buf.WriteString(fmt.Sprintf("%s%v\n", prefix, node))
	if n > 0 {
		if end {
			prefix = prefix[:len(prefix)-3] + "   "
		} else {
			prefix = prefix[:len(prefix)-3] + "|  "
		}
	}
	i := 0
	for ; i < len(node.GetChildren())-1; i++ {
		prePrintTree(buf, prefix, node.GetChildren()[i], n+1, false)
	}
	if i < len(node.GetChildren()) && len(node.GetChildren()) > 0 {
		prePrintTree(buf, prefix, node.GetChildren()[i], n+1, true)
	}
}

// TreeNodeGraph 多叉树
type TreeNodeGraph interface {
	Children() []TreeNodeGraph
	GetValue() interface{}
	GetAttribute() string
}

// todo TreeSvg 创建多叉树svg图片
func TreeSvg(root TreeNodeGraph, fileName string) error {
	return nil
}
