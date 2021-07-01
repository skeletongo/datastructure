package common

import "fmt"

type INode interface {
	GetLeftNode() INode
	GetRightNode() INode
	GetValue() interface{}
}

type INodeKey interface {
	GetLeftNode() INodeKey
	GetRightNode() INodeKey
	GetKey() interface{}
	GetValue() interface{}
}

func getLeftNode(v interface{}) interface{} {
	switch val := v.(type) {
	case INode:
		return val.GetLeftNode()
	case INodeKey:
		return val.GetLeftNode()
	default:
		panic("invalid node")
	}
}

func getRightNode(v interface{}) interface{} {
	switch val := v.(type) {
	case INode:
		return val.GetRightNode()
	case INodeKey:
		return val.GetRightNode()
	default:
		panic("invalid node")
	}
}

func getKey(v interface{}) interface{} {
	switch val := v.(type) {
	case INode:
		return val.GetValue()
	case INodeKey:
		return val.GetKey()
	default:
		panic("invalid node")
	}
}

func getValue(v interface{}) interface{} {
	switch val := v.(type) {
	case INode:
		return nil
	case INodeKey:
		return val.GetValue()
	default:
		panic("invalid node")
	}
}

func getString(v interface{}) string {
	switch val := v.(type) {
	case INode:
		return fmt.Sprint(val.GetValue())
	case INodeKey:
		return fmt.Sprintf("%v: %v", val.GetKey(), val.GetValue())
	default:
		panic("invalid node")
	}
}
