package hashtable

import (
	"container/list"
)

type TableList struct {
	size      int // 元素数量
	index     int // 取模用到的素数在 capacity 中的下标
	m         int
	tableList []*list.List
}

func NewTableList() *TableList {
	ret := new(TableList)
	ret.m = capacity[0]
	ret.tableList = make([]*list.List, ret.m)
	for i := 0; i < len(ret.tableList); i++ {
		ret.tableList[i] = list.New()
	}
	return ret
}

func (t *TableList) GetSize() int {
	return t.size
}

func (t *TableList) IsEmpty() bool {
	return t.size == 0
}

func (t *TableList) resize() {
	t.m = capacity[t.index]
	newTableList := make([]*list.List, t.m)
	for i := 0; i < len(newTableList); i++ {
		newTableList[i] = list.New()
	}
	for _, v := range t.tableList {
		for e := v.Front(); e != nil; e = e.Next() {
			node := e.Value.([]interface{})
			code := hash(node[0], t.m)
			newTableList[code].PushBack(node)
		}
	}
	t.tableList = newTableList
}

func (t *TableList) Set(key, value interface{}) {
	l := t.tableList[hash(key, t.m)]
	for e := l.Front(); e != nil; e = e.Next() {
		node := e.Value.([]interface{})
		if node[0] == key {
			node[1] = value
			return
		}
	}
	l.PushBack([]interface{}{key, value})
	t.size++
	// 扩容
	if t.size > t.m*upperTol && t.index+1 < len(capacity) {
		t.index++
		t.resize()
	}
}

func (t *TableList) Remove(key interface{}) {
	l := t.tableList[hash(key, t.m)]
	for e := l.Front(); e != nil; e = e.Next() {
		node := e.Value.([]interface{})
		if node[0] == key {
			l.Remove(e)
			t.size--
			// 缩容
			if t.size < t.m*lowerTol && t.index-1 >= 0 {
				t.index--
				t.resize()
			}
			return
		}
	}
	return
}

func (t *TableList) Contains(key interface{}) bool {
	l := t.tableList[hash(key, t.m)]
	for e := l.Front(); e != nil; e = e.Next() {
		node := e.Value.([]interface{})
		if node[0] == key {
			return true
		}
	}
	return false
}

func (t *TableList) Get(key interface{}) interface{} {
	l := t.tableList[hash(key, t.m)]
	for e := l.Front(); e != nil; e = e.Next() {
		node := e.Value.([]interface{})
		if node[0] == key {
			return node[1]
		}
	}
	return nil
}
