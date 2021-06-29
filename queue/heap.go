package queue

import (
	"container/heap"
)

/*
使用标准库的堆，实现优先队列；标准库已经给了示例代码，也可以参考标准库
*/

// Item 队列元素
type Item struct {
	Value interface{} // 元素值
	index int         // 只能 PriorityQueue 修改
}

// GetIndex 查看元素唯一索引
func (i *Item) GetIndex() int {
	return i.index
}

// PriorityQueue 优先队列
type PriorityQueue struct {
	Items   []*Item
	Compare func(a, b interface{}) int
}

// NewPriorityQueue 创建优先队列
// 参数 compare 为自定义元素大小比较函数
// 大小比较函数 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b
func NewPriorityQueue(compare func(a, b interface{}) int) *PriorityQueue {
	return &PriorityQueue{Compare: compare}
}

func (q *PriorityQueue) Len() int {
	return len(q.Items)
}

func (q *PriorityQueue) Less(i, j int) bool {
	return q.Compare(q.Items[i].Value, q.Items[j].Value) < 0
}

func (q *PriorityQueue) Swap(i, j int) {
	q.Items[i], q.Items[j] = q.Items[j], q.Items[i]
	q.Items[i].index = i
	q.Items[j].index = j
}

func (q *PriorityQueue) Push(item interface{}) {
	if item == nil {
		return
	}
	x := item.(*Item)
	x.index = len(q.Items)
	q.Items = append(q.Items, x)
}

func (q *PriorityQueue) Pop() interface{} {
	n := len(q.Items)
	item := q.Items[n-1]
	q.Items[n-1] = nil // avoid memory leak
	item.index = -1    // for safety
	q.Items = q.Items[0 : n-1]
	return item
}

// Init 堆化
func (q *PriorityQueue) Init() {
	heap.Init(q)
}

// Enqueue 入队
// v 必须是*Item类型
func (q *PriorityQueue) Enqueue(v interface{}) {
	heap.Push(q, v.(*Item))
}

// Dequeue 出队
// 返回值是*Item类型
func (q *PriorityQueue) Dequeue() interface{} {
	if q.Len() == 0 {
		return nil
	}
	return heap.Pop(q)
}

func (q *PriorityQueue) Update(item *Item) {
	heap.Fix(q, item.index)
}

// HeapInit 堆化
func HeapInit(q *PriorityQueue) {
	heap.Init(q)
}

// HeapPush 入队
func HeapPush(q *PriorityQueue, item *Item) {
	heap.Push(q, item)
}

// HeapPop 出队
func HeapPop(q *PriorityQueue) *Item {
	if q.Len() == 0 {
		return nil
	}
	return heap.Pop(q).(*Item)
}

// HeapUpdate 元素值修改后更新优先队列(维护堆的性质)
func HeapUpdate(q *PriorityQueue, item *Item) {
	heap.Fix(q, item.index)
}
