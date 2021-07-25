package common

import (
	"bytes"
	"fmt"
	"io"
)

// BSTNode 二分查找树
type BSTNode interface {
	GetLeftNode() BSTNode
	GetRightNode() BSTNode
	GetValue() interface{}
}

// PrePrintBST 前序打印二叉树
func PrePrintBST(n BSTNode) string {
	if IsNil(n) {
		return ""
	}
	buf := &bytes.Buffer{}
	prePrintBST(buf, "", n, 0, false, false)
	return buf.String()
}

func prePrintBST(buf *bytes.Buffer, prefix string, node BSTNode, n int, end, isLeft bool) {
	// 接口值包含两部分，所包含的动态类型和动态类型的值
	// 只有两部分都为nil，接口值才等于nil
	if IsNil(node) {
		return
	}
	if n > 0 {
		if end && !isLeft {
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
	prePrintBST(buf, prefix, node.GetLeftNode(), n+1, IsNil(node.GetRightNode()), true)
	prePrintBST(buf, prefix, node.GetRightNode(), n+1, true, false)
}

// PrePrintBSTSlice 前序打印用数组实现的二叉树
// arr 二叉树切片
func PrePrintBSTSlice(arr []interface{}) string {
	buf := &bytes.Buffer{}
	prePrintBSTSlice(buf, "", arr, 0, 0, false, false)
	return buf.String()
}

func prePrintBSTSlice(buf *bytes.Buffer, prefix string, arr []interface{}, i, n int, end, isLeft bool) {
	if i >= len(arr) {
		return
	}
	if n > 0 {
		if end && !isLeft {
			prefix += "`--"
		} else {
			prefix += "|--"
		}
	}
	buf.WriteString(fmt.Sprintf("%s(%v\n", prefix, arr[i]))
	if n > 0 {
		if end {
			prefix = prefix[:len(prefix)-3] + "   "
		} else {
			prefix = prefix[:len(prefix)-3] + "|  "
		}
	}
	leftIndex := 2*i + 1
	rightIndex := 2*i + 2
	prePrintBSTSlice(buf, prefix, arr, leftIndex, n+1, rightIndex >= len(arr), true)
	prePrintBSTSlice(buf, prefix, arr, rightIndex, n+1, true, false)
}

// BSTNodeGraph 二分查找树
type BSTNodeGraph interface {
	LeftNode() BSTNodeGraph
	RightNode() BSTNodeGraph
	GetValue() interface{}
	GetAttribute() string
}

func bstPostOrder(w io.Writer, n BSTNodeGraph, m map[BSTNodeGraph]int, i *int) {
	if IsNil(n) {
		return
	}

	bstPostOrder(w, n.LeftNode(), m, i)
	bstPostOrder(w, n.RightNode(), m, i)

	w.Write([]byte(fmt.Sprintf("%v[label = \"<f0> |<f1> %v|<f2> \"%v]\n", *i, n, n.GetAttribute())))

	if !IsNil(n.LeftNode()) {
		w.Write([]byte(fmt.Sprintf("\"%v\":f0 -> \"%v\":f1\n", *i, m[n.LeftNode()])))
	}
	if !IsNil(n.RightNode()) {
		w.Write([]byte(fmt.Sprintf("\"%v\":f2 -> \"%v\":f1\n", *i, m[n.RightNode()])))
	}
	m[n] = *i
	*i++
}

func NewBSTDot(w io.Writer, root BSTNodeGraph) (err error) {
	if IsNil(root) {
		return
	}
	if _, err = w.Write([]byte("digraph G {\nnode [shape = record,height=.1]\n")); err != nil {
		return
	}

	i := 0
	m := map[BSTNodeGraph]int{}
	bstPostOrder(w, root, m, &i)

	_, err = w.Write([]byte(`}`))
	return
}

// BSTSvg 创建二叉查找树svg图片
func BSTSvg(root BSTNodeGraph, filename string) error {
	return NewImg(filename, root, func(w io.Writer, root interface{}) error {
		return NewBSTDot(w, root.(BSTNodeGraph))
	})
}
