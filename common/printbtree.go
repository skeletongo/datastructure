package common

import (
	"fmt"
	"io"
)

// BTreeNodeGraph B树
type BTreeNodeGraph interface {
	Children() []BTreeNodeGraph
	Values() []interface{}
}

func btreePostOrder(w io.Writer, n BTreeNodeGraph, m map[BTreeNodeGraph]int, i *int) {
	if IsNil(n) {
		return
	}

	for _, v := range n.Children() {
		btreePostOrder(w, v, m, i)
	}

	w.Write([]byte(fmt.Sprintf("%v[label = \"", *i)))
	w.Write([]byte(fmt.Sprintf("<c0> |<v0> %v", n.Values()[0])))
	for k, v := range n.Values()[1:] {
		w.Write([]byte(fmt.Sprintf("|<c%d> ", k+1)))
		w.Write([]byte(fmt.Sprintf("|<v%d> %v", k+1, v)))
	}
	w.Write([]byte(fmt.Sprintf("|<c%d> \"]\n", len(n.Values()))))

	for k, v := range n.Children() {
		p := "v"
		if len(v.Values())%2 == 0 {
			p = "c"
		}
		w.Write([]byte(fmt.Sprintf("\"%v\":c%d -> \"%v\":%s%d\n",
			*i,
			k,
			m[v],
			p,
			len(v.Values())/2,
		)))
	}

	m[n] = *i
	*i++
}

func NewBTreeDot(w io.Writer, root BTreeNodeGraph) (err error) {
	if IsNil(root) {
		return
	}
	if _, err = w.Write([]byte("digraph G {\nnode [shape = record,height=.1]\n")); err != nil {
		return
	}

	i := 0
	m := map[BTreeNodeGraph]int{}
	btreePostOrder(w, root, m, &i)

	_, err = w.Write([]byte(`}`))
	return
}

// BTreeSvg 创建B树svg图片
func BTreeSvg(root BTreeNodeGraph, filename string) error {
	return NewImg(filename, root, func(w io.Writer, root interface{}) error {
		return NewBTreeDot(w, root.(BTreeNodeGraph))
	})
}
