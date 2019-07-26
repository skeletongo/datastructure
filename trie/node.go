package dataStructure

type node struct {
	isWord bool
	next   map[string]*node
}

func newNode() *node {
	return &node{
		isWord: false,
		next:   make(map[string]*node),
	}
}
