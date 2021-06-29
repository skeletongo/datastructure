package heap

import "github.com/skeletongo/datastructure/common"

// IndexHeap 索引堆
type IndexHeap struct {
	// data 保存添加的元素
	data []interface{}

	// indexes 根据元素值对元素的data索引进行堆化
	indexes []int

	// reverse 元素data索引在indexes中的位置; reverse[indexes[i]] = i
	reverse []int

	arr []int // 可用data位置

	// 元素比较方法
	Compare func(a, b interface{}) int
}

// NewIndexHeap 创建索引堆
// 参数 compare 为自定义元素大小比较函数
// 大小比较函数 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b
func NewIndexHeap(compare func(a, b interface{}) int) *IndexHeap {
	return &IndexHeap{Compare: compare}
}

// HeapifyIndexHeap 使用切片创建最大堆
// 从最后一个有子节点的节点开始做元素下沉操作
// 参数 compare 为自定义元素大小比较函数
// 大小比较函数 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b
func HeapifyIndexHeap(arr []interface{}, compare func(a, b interface{}) int) *IndexHeap {
	h := NewIndexHeap(compare)
	h.heapify(arr)
	return h
}

func (h *IndexHeap) GetSize() int {
	return len(h.indexes)
}

func (h *IndexHeap) IsEmpty() bool {
	return len(h.indexes) == 0
}

func (h *IndexHeap) shiftUp(i int) {
	j := (i - 1) / 2
	e := h.indexes[i]
	for i > 0 && h.Compare(h.data[e], h.data[h.indexes[j]]) > 0 {
		h.indexes[i] = h.indexes[j]
		h.reverse[h.indexes[i]] = i
		i = j
		j = (i - 1) / 2
	}
	h.indexes[i] = e
	h.reverse[e] = i
}

func (h *IndexHeap) shiftDown(i int) {
	n := len(h.indexes)
	e := h.indexes[i]

	for {
		j := 2*i + 1
		if j >= n || j < 0 { // j < 0,当 2*i+1 整数溢出时
			break
		}

		if ri := j + 1; ri < n && h.Compare(h.data[h.indexes[j]], h.data[h.indexes[ri]]) < 0 {
			j = ri
		}
		if h.Compare(h.data[e], h.data[h.indexes[j]]) >= 0 {
			break
		}
		h.indexes[i] = h.indexes[j]
		h.reverse[h.indexes[i]] = i
		i = j
	}
	h.indexes[i] = e
	h.reverse[e] = i
}

func (h *IndexHeap) contains(index int) bool {
	if index < 0 || index >= len(h.data) {
		return false
	}
	return h.reverse[index] != -1
}

func (h *IndexHeap) Get(index int) interface{} {
	if !h.contains(index) {
		panic("Invalid index")
	}
	return h.data[index]
}

func (h *IndexHeap) Change(index int, v interface{}) {
	if !h.contains(index) {
		panic("Invalid index")
	}

	h.data[index] = v
	j := h.reverse[index]
	if len(h.indexes) > 1 {
		h.shiftUp(j)
		h.shiftDown(j)
	}
}

func (h *IndexHeap) Add(v interface{}) int {
	// 获取可用位置
	index := -1
	if len(h.arr) > 0 {
		index = h.arr[len(h.arr)-1]
		h.arr = h.arr[:len(h.arr)-1]
	}

	if index == -1 { // 追加元素
		h.data = append(h.data, v)
		index = len(h.data) - 1
		h.reverse = append(h.reverse, len(h.indexes))
	} else { // 覆盖已取出元素的值
		h.data[index] = v
		h.reverse[index] = len(h.indexes)
	}

	h.indexes = append(h.indexes, index)
	if len(h.indexes) > 1 {
		h.shiftUp(len(h.indexes) - 1)
	}
	return index
}

func (h *IndexHeap) ExtractMax() interface{} {
	if len(h.indexes) == 0 {
		panic("no data")
	}
	v := h.data[h.indexes[0]]
	h.arr = append(h.arr, h.indexes[0]) // 记录已取出元素的data下标

	h.indexes[0], h.indexes[len(h.indexes)-1] = h.indexes[len(h.indexes)-1], h.indexes[0]
	h.reverse[h.indexes[0]] = 0
	h.reverse[h.indexes[len(h.indexes)-1]] = -1
	h.indexes = h.indexes[:len(h.indexes)-1]
	if len(h.indexes) > 1 {
		h.shiftDown(0)
	}
	return v
}

func (h *IndexHeap) heapify(arr []interface{}) {
	for k, v := range arr {
		h.data = append(h.data, v)
		h.indexes = append(h.indexes, k)
		h.reverse = append(h.reverse, k)
	}
	if len(arr) < 2 {
		return
	}
	for i := (len(arr) - 2) / 2; i >= 0; i-- {
		h.shiftDown(i)
	}
}

func (h *IndexHeap) String() string {
	return common.PrePrintBSTSlice(h.data)
}
