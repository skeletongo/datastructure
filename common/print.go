package common

import (
	"bytes"
	"fmt"
	"reflect"
)

type INode interface {
	GetLeftNode() INode
	GetRightNode() INode
	GetValue() interface{}
}

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
func PrePrint(n INode) string {
	buf := &bytes.Buffer{}
	prePrint(buf, "", n, 0, false, false)
	return buf.String()
}

func prePrint(buf *bytes.Buffer, prefix string, node INode, n int, end, isLeft bool) {
	// 接口值包含两部分，所包含的动态类型和动态类型的值
	// 只有两部分都为nil，接口值才等于nil
	if reflect.ValueOf(node).IsNil() {
		return
	}
	if n > 0 {
		if end && !isLeft {
			prefix += "`--"
		} else {
			prefix += "|--"
		}
	}
	buf.WriteString(fmt.Sprintf("%s%v\n", prefix, node.GetValue()))
	if n > 0 {
		if end {
			prefix = prefix[:len(prefix)-3] + "   "
		} else {
			prefix = prefix[:len(prefix)-3] + "|  "
		}
	}
	prePrint(buf, prefix, node.GetLeftNode(), n+1, reflect.ValueOf(node.GetRightNode()).IsNil(), true)
	prePrint(buf, prefix, node.GetRightNode(), n+1, true, false)
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
