package common

import "reflect"

func PreOrder(n interface{}, f func(key, value interface{})) {
	if n == nil || reflect.ValueOf(n).IsNil() {
		return
	}

	f(getKey(n), getValue(n))
	PreOrder(getLeftNode(n), f)
	PreOrder(getRightNode(n), f)
}

func InOrder(n interface{}, f func(key, value interface{})) {
	if n == nil || reflect.ValueOf(n).IsNil() {
		return
	}

	InOrder(getLeftNode(n), f)
	f(getKey(n), getValue(n))
	InOrder(getRightNode(n), f)
}

func PostOrder(n interface{}, f func(key, value interface{})) {
	if n == nil || reflect.ValueOf(n).IsNil() {
		return
	}

	PostOrder(getLeftNode(n), f)
	PostOrder(getRightNode(n), f)
	f(getKey(n), getValue(n))
}
