package dataStructure

type Node struct {
	Num   int
	Left  *Node
	Right *Node
}

func NewNode(n int) *Node {
	return &Node{Num: n}
}

type TrieNode struct {
	IsWord bool
	Next   map[string]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		IsWord: false,
		Next:   make(map[string]*TrieNode),
	}
}
