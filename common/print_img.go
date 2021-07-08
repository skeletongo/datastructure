package common

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
)

/*
生成二叉树图
*/

type ITreeNode interface {
	GetLeftTreeNode() ITreeNode
	GetRightTreeNode() ITreeNode
	IsWrite() bool
	Write()
	GetIndex() int
	SetIndex(index int)
	GetColor() string
	Reset()
}

type TreeNode struct {
	Index   int
	isWrite bool
}

func (t *TreeNode) GetIndex() int {
	return t.Index
}

func (t *TreeNode) SetIndex(index int) {
	t.Index = index
}

func (t *TreeNode) IsWrite() bool {
	return t.isWrite
}

func (t *TreeNode) Write() {
	t.isWrite = true
}

func (t *TreeNode) Reset() {
	t.Index = 0
	t.isWrite = false
}

func StyleByColor(color string) string {
	switch color {
	case "red":
		return ",color=red,style=filled,fontcolor=white"
	case "black":
		return ",color=black,style=filled,fontcolor=white"
	default:
		return ""
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func NewDir(path string) error {
	exist, err := PathExists(path)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}
	if err := os.Mkdir(path, 0666); err != nil {
		return err
	}
	return nil
}

func PrintTree(root ITreeNode, fileName string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	dirSvg := path.Join(dir, "img")
	dirDot := path.Join(dirSvg, "dot")
	if err = NewDir(dirSvg); err != nil {
		return err
	}
	if err = NewDir(dirDot); err != nil {
		return err
	}

	if IsNil(root) {
		return nil
	}

	svgFile := fmt.Sprintf("%s.svg", path.Join(dirSvg, fileName))
	dotFile := fmt.Sprintf("%s.dot", path.Join(dirDot, fileName))
	fw, err := os.Create(dotFile)
	if err != nil {
		return err
	}
	fw.WriteString(`digraph G {
    graph [nodesep=0.1]
    node [shape=circle]
    edge [arrowhead=vee]
`)
	if !IsNil(root.GetLeftTreeNode()) || !IsNil(root.GetRightTreeNode()) {
		root.Write()
		fmt.Fprintf(fw, "    %d [group=%d,label=\"%v\"%s]\n", root.GetIndex(), root.GetIndex(), root, StyleByColor(root.GetColor()))
	}
	printNode(fw, root)
	fw.WriteString("}")
	fw.Close()

	cmd := exec.Command("dot", dotFile, "-Tsvg", "-o", svgFile)
	if err = cmd.Run(); err != nil {
		return err
	}
	return nil
}

func printNode(fw io.Writer, root ITreeNode) {
	if !root.IsWrite() {
		fmt.Fprintf(fw, "    %d [label=\"%v\"%s]\n", root.GetIndex(), root, StyleByColor(root.GetColor()))
	}
	target, distance := 0, 0
	if !IsNil(root.GetLeftTreeNode()) {
		leftMax := root
		leftDistance := 1
		for !IsNil(leftMax.GetRightTreeNode()) {
			leftMax = leftMax.GetRightTreeNode()
			leftDistance++
		}
		// 找到root节点的root.left往下最右边的节点
		target = leftMax.GetIndex()
		distance = leftDistance
		if !IsNil(root.GetLeftTreeNode().GetLeftTreeNode()) || !IsNil(root.GetLeftTreeNode().GetRightTreeNode()) {
			root.GetLeftTreeNode().Write() // 生成该节点值
			fmt.Fprintf(fw, "    %d [group=%d,label=\"%v\"%s]\n", root.GetLeftTreeNode().GetIndex(), root.GetLeftTreeNode().GetIndex(), root.GetLeftTreeNode(), StyleByColor(root.GetLeftTreeNode().GetColor()))
		}
		// 生成root指向root.left的关系
		fmt.Fprintf(fw, "    %d -> %d\n", root.GetIndex(), root.GetLeftTreeNode().GetIndex())
		printNode(fw, root.GetLeftTreeNode())
	}

	if !IsNil(root.GetLeftTreeNode()) || !IsNil(root.GetRightTreeNode()) {
		// 弄一个中间节点,隐藏起来,主要是让布局更美观
		fmt.Fprintf(fw, "    _%d [group=%d,label=\"\",width=0,style=invis]\n", root.GetIndex(), root.GetIndex())
		fmt.Fprintf(fw, "    %d -> _%d [style=invis]\n", root.GetIndex(), root.GetIndex())
	}

	if !IsNil(root.GetRightTreeNode()) {
		rightMin := root.GetRightTreeNode()
		rightDistance := 1
		for !IsNil(rightMin.GetLeftTreeNode()) {
			rightMin = rightMin.GetLeftTreeNode()
			rightDistance++
		}
		// 找到root节点的root.Right往下最左边的节点
		if rightDistance <= distance {
			target = rightMin.GetIndex()
			distance = rightDistance
		}
		if !IsNil(root.GetRightTreeNode().GetLeftTreeNode()) || !IsNil(root.GetRightTreeNode().GetRightTreeNode()) {
			root.GetRightTreeNode().Write() // 生成该节点值
			fmt.Fprintf(fw, "    %d [group=%d,label=\"%v\"%s]\n", root.GetRightTreeNode().GetIndex(), root.GetRightTreeNode().GetIndex(), root.GetRightTreeNode(), StyleByColor(root.GetRightTreeNode().GetColor()))
		}
		// 生成root指向root.Right的关系
		fmt.Fprintf(fw, "    %d -> %d\n", root.GetIndex(), root.GetRightTreeNode().GetIndex())
		printNode(fw, root.GetRightTreeNode())
	}

	// 一个节点对应的占位节点应该与该节点的左子树的最大节点和右子树的最小节点中距离较近的那一个处于同一层
	if distance > 1 && target != 0 {
		fmt.Fprintf(fw, "    {rank=same;_%d;%d}\n", root.GetIndex(), target)
	}
}
