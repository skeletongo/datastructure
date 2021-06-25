// 最大堆
package heap

import "dataStructure/common"

/*
 二叉堆 添加和取出元素的时间复杂度 O(logN)
 使用完全二叉树实现，利用数组存储
*/

// ArrayHeap 最大堆
type ArrayHeap struct {
	array []interface{}
	f     func(a, b interface{}) int
}

// NewArrayHeap 创建堆
// 参数 f 为自定义元素大小比较函数
// 大小比较函数 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b
func NewArrayHeap(f func(a, b interface{}) int) *ArrayHeap {
	return &ArrayHeap{f: f}
}

// GetSize 元素数量
func (h *ArrayHeap) GetSize() int {
	return len(h.array)
}

// IsEmpty 是否为空
func (h *ArrayHeap) IsEmpty() bool {
	return h.GetSize() == 0
}

// parent 查询父节点索引
func parent(i int) int {
	if i == 0 {
		panic("parent: no parent")
	}
	return (i - 1) / 2
}

// leftChild 查询左节点索引
//func leftChild(i int) int {
//	return 2*i + 1
//}

// rightChild 查询右节点索引
//func rightChild(i int) int {
//	return 2*i + 2
//}

// findMax 查询最大元素
func (h *ArrayHeap) findMax() interface{} {
	if h.GetSize() == 0 {
		panic("findMax: no data")
	}
	return h.array[0]
}

// siftUp 元素上浮
func (h *ArrayHeap) siftUp(i int) {
	for i > 0 && h.f(h.array[i], h.array[parent(i)]) > 0 {
		j := parent(i)
		h.array[i], h.array[j] = h.array[j], h.array[i]
		i = j
	}
}

// Add 添加元素
// 时间复杂度 O(logN)
func (h *ArrayHeap) Add(e interface{}) {
	h.array = append(h.array, e)
	h.siftUp(h.GetSize() - 1)
}

// siftDown 元素下沉
func (h *ArrayHeap) siftDown(i int) bool {
	i0 := i
	n := h.GetSize()
	for {
		j := 2*i + 1
		if j >= n || j < 0 { // j < 0,当 2*i+1 整数溢出时
			break
		}
		if ri := j + 1; ri < n && h.f(h.array[j], h.array[ri]) < 0 {
			j = ri
		}
		if h.f(h.array[i], h.array[j]) >= 0 {
			break
		}
		h.array[i], h.array[j] = h.array[j], h.array[i]
		i = j
	}
	return i > i0
}

// ExtractMax 取出元素
// 时间复杂度 O(logN)
func (h *ArrayHeap) ExtractMax() interface{} {
	t := h.findMax()
	h.array[0] = h.array[h.GetSize()-1]
	h.array = h.array[:h.GetSize()-1]
	h.siftDown(0)
	return t
}

// Replace 取出并添加一个元素
// 时间复杂度 O(logN)
func (h *ArrayHeap) Replace(e interface{}) interface{} {
	t := h.findMax()
	h.array[0] = e
	h.siftDown(0)
	return t
}

// Heapify 使用切片创建最大堆
// 时间复杂度 O(N)
// 方式：从最后一个有子节点的节点开始做元素下沉操作
func (h *ArrayHeap) Heapify(arr []interface{}) {
	for _, v := range arr {
		h.array = append(h.array, v)
	}
	if len(arr) <= 1 {
		return
	}
	i := parent(h.GetSize() - 1)
	for ; i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *ArrayHeap) String() string {
	return common.PrePrintBSTSlice(h.array)
}
