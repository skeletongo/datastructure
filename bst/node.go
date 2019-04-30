package bst

type node struct {
	e     interface{}
	left  *node
	right *node
}

func newNode(e interface{}) *node {
	return &node{e: e}
}
