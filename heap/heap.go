// 堆
package heap

import "dataStructure/common"

/*
 二叉堆 添加和取出元素的时间复杂度 O(logN)
 使用完全二叉树实现，利用数组存储
*/

// Heap 堆
type Heap struct {
	array   []interface{}
	Compare func(a, b interface{}) int
}

// New 创建最大堆
// 参数 compare 为自定义元素大小比较函数
// 大小比较函数 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b
func New(compare func(a, b interface{}) int) *Heap {
	return &Heap{Compare: compare}
}

// Heapify 使用切片创建最大堆
// 从最后一个有子节点的节点开始做元素下沉操作
// 参数 compare 为自定义元素大小比较函数
// 大小比较函数 返回值：
// 负数	表示	a<b
// 0	表示	a=b
// 正数	表示	a>b
func Heapify(arr []interface{}, f func(a, b interface{}) int) *Heap {
	h := New(f)
	h.heapify(arr)
	return h
}

// GetSize 元素数量
func (h *Heap) GetSize() int {
	return len(h.array)
}

// IsEmpty 是否为空
func (h *Heap) IsEmpty() bool {
	return len(h.array) == 0
}

// siftUp 元素上浮
func (h *Heap) siftUp(i int) {
	j := (i - 1) / 2
	e := h.array[i]
	for i > 0 && h.Compare(e, h.array[j]) > 0 {
		h.array[i] = h.array[j]
		i = j
		j = (i - 1) / 2
	}
	h.array[i] = e
}

// siftDown 元素下沉
func (h *Heap) siftDown(i int) bool {
	i0 := i
	n := len(h.array)
	e := h.array[i]
	for {
		j := 2*i + 1
		if j >= n || j < 0 { // j < 0,当 2*i+1 整数溢出时
			break
		}
		if ri := j + 1; ri < n && h.Compare(h.array[j], h.array[ri]) < 0 {
			j = ri
		}
		if h.Compare(e, h.array[j]) >= 0 {
			break
		}
		h.array[i] = h.array[j]
		i = j
	}
	h.array[i] = e
	return i > i0
}

// Add 添加元素
func (h *Heap) Add(v interface{}) {
	h.array = append(h.array, v)
	if len(h.array) > 1 {
		h.siftUp(len(h.array) - 1)
	}
}

// ExtractMax 取出元素
func (h *Heap) ExtractMax() interface{} {
	if len(h.array) == 0 {
		return nil
	}
	t := h.array[0]
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	if len(h.array) > 1 {
		h.siftDown(0)
	}
	return t
}

// Replace 取出并添加一个元素
func (h *Heap) Replace(e interface{}) interface{} {
	if len(h.array) == 0 {
		h.Add(e)
		return nil
	}
	t := h.array[0]
	h.array[0] = e
	if len(h.array) > 1 {
		h.siftDown(0)
	}
	return t
}

func (h *Heap) heapify(arr []interface{}) {
	for _, v := range arr {
		h.array = append(h.array, v)
	}
	if len(arr) < 2 {
		return
	}
	i := (len(h.array) - 2) / 2
	for ; i >= 0; i-- {
		h.siftDown(i)
	}
}

func (h *Heap) String() string {
	return common.PrePrintBSTSlice(h.array)
}
