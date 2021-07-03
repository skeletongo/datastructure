package common

import (
	"bytes"
	"fmt"
	"reflect"
)

// 前序打印
/////////////////
//5            //
//|--3         //
//|  |--2      //
//|  |  |--1   //
//|  `--4      //
//`--8         //
//   |--6      //
//      `--7   //
/////////////////

// PrePrint 前序打印二叉树
// n 二叉树节点
func PrePrint(n interface{}) string {
	if n == nil || reflect.ValueOf(n).IsNil() {
		return ""
	}
	buf := &bytes.Buffer{}
	prePrint(buf, "", n, 0, false, false)
	return buf.String()
}

func prePrint(buf *bytes.Buffer, prefix string, node interface{}, n int, end, isLeft bool) {
	// 接口值包含两部分，所包含的动态类型和动态类型的值
	// 只有两部分都为nil，接口值才等于nil
	if node == nil || reflect.ValueOf(node).IsNil() {
		return
	}
	if n > 0 {
		if end && !isLeft {
			prefix += "`--"
		} else {
			prefix += "|--"
		}
	}
	buf.WriteString(fmt.Sprintf("%s(%v\n", prefix, getString(node)))
	if n > 0 {
		if end {
			prefix = prefix[:len(prefix)-3] + "   "
		} else {
			prefix = prefix[:len(prefix)-3] + "|  "
		}
	}
	prePrint(buf, prefix, getLeftNode(node), n+1, getRightNode(node) == nil, true)
	prePrint(buf, prefix, getRightNode(node), n+1, true, false)
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

// todo 层序打印
/////////////////
//       5     //
//     /   \   //
//    3     8  //
//   / \   /   //
//  2   4 6    //
// /       \   //
// 1        7  //
/////////////////
//func LevelPrint(n INode) string {
//
//}
