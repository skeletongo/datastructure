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

type AVLNode struct {
	Key int // 可比较的类型，只是演示用，没有做过多考虑

	Height int
	Left   *AVLNode
	Right  *AVLNode
}

func NewAVLNode(key int) *AVLNode {
	return &AVLNode{
		Key:    key,
		Height: 1,
	}
}
