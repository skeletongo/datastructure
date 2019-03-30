package dataStructure

type Node struct {
	Num   int
	Left  *Node
	Right *Node
}

func NewNode(n int) *Node {
	return &Node{Num: n}
}
