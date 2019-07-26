package common

type Unioner interface {
	GetSize() int
	IsConnected(p, q int) bool
	UnionElements(p, q int)
}
