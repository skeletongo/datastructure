package dataStructure

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
