package heap

import "dataStructure/array"

/*
 二叉堆 添加和取出元素的时间复杂度 O(logN)
 使用完全二叉树实现，利用数组存储
*/

// 最大堆
type Heap struct {
	array *array.Array
	f     func(a, b interface{}) int
}

func (h *Heap) GetSize() int {
	return h.array.GetSize()
}

func (h *Heap) IsEmpty() bool {
	return h.array.GetSize() == 0
}

// 添加元素
func (h *Heap) Add(e interface{}) {
	h.array.AddLast(e)
	h.siftUp(h.GetSize() - 1)
}

// 取出元素
func (h *Heap) ExtractMax() interface{} {
	e := h.findMax()
	h.array.Set(0, h.array.Get(h.GetSize()-1))
	h.array.RemoveLast()
	h.siftDown(0)
	return e
}

// 取出并添加一个元素
func (h *Heap) Replace(e interface{}) interface{} {
	res := h.findMax()
	h.array.Set(0, e)
	h.siftDown(0)
	return res
}

// 使用堆存储数据
// 方式：从最后一个有子节点的节点开始做元素下沉操作
func (h *Heap) Heapify(arr []interface{}) {
	for _, v := range arr {
		h.array.AddLast(v)
	}
	k := h.parent(h.GetSize() - 1)
	for ; k >= 0; k-- {
		h.siftDown(k)
	}
}

// 返回最大值
func (h *Heap) findMax() interface{} {
	if h.GetSize() == 0 {
		panic("no data")
	}
	return h.array.Get(0)
}

// 查询父节点索引
func (h *Heap) parent(k int) int {
	if k == 0 {
		panic("no parent")
	}
	return (k - 1) / 2
}

// 查询左节点索引
func (h *Heap) leftChild(k int) int {
	return 2*k + 1
}

// 查询右节点索引
func (h *Heap) rightChild(k int) int {
	return 2*k + 2
}

// 元素上浮
func (h *Heap) siftUp(k int) {
	for ; k > 0 && h.f(h.array.Get(h.parent(k)), h.array.Get(k)) < 0; k = h.parent(k) {
		// 父节点小于子节点，需要交换位置
		h.array.Swap(h.parent(k), k)
	}
}

// 元素下沉
func (h *Heap) siftDown(k int) {
	for h.leftChild(k) < h.GetSize() { // 至少有左节点
		j := h.leftChild(k)
		// 找出子节点中的最大值对应的索引
		if j+1 < h.GetSize() && h.f(h.array.Get(j), h.array.Get(j+1)) < 0 {
			j++
		}
		// 判度是否需要交换位置
		if h.f(h.array.Get(j), h.array.Get(k)) <= 0 {
			break
		}
		// 交换
		h.array.Swap(k, j)
		k = j
	}
}

// 创建堆
// 参数 f 为自定义元素大小比较函数
// 大小比较函数 返回值：
// -1	表示	a<b
// 0	表示	a=b
// 1	表示	a>b
//func f(a,b interface{}) int { }
func NewHeap(f func(a, b interface{}) int) *Heap {
	return &Heap{array: array.NewArray(), f: f}
}
