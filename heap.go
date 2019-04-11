package dataStructure

/*
 二叉堆 添加和取出元素的时间复杂度 O(logN)
 使用完全二叉树实现，利用数组存储
*/

// 最大堆
type MaxHeap struct {
	array []int
}

func (m *MaxHeap) GetSize() int {
	return len(m.array)
}

func (m *MaxHeap) IsEmpty() bool {
	return len(m.array) == 0
}

func (m *MaxHeap) parent(k int) int {
	if k == 0 {
		panic("param error")
	}
	return (k - 1) / 2
}

func (m *MaxHeap) leftChild(k int) int {
	return k*2 + 1
}

func (m *MaxHeap) rightChild(k int) int {
	return k*2 + 2
}

func (m *MaxHeap) findMax() int {
	if len(m.array) == 0 {
		panic("no data")
	}
	return m.array[0]
}

func (m *MaxHeap) Add(e int) {
	// 新添加元素，添加到最后位置
	m.array = append(m.array, e)
	// 维护堆结构
	m.siftUp(len(m.array) - 1)
}

// 元素上浮
func (m *MaxHeap) siftUp(k int) {
	for k > 0 && m.array[k] > m.array[m.parent(k)] {
		m.array[k], m.array[m.parent(k)] = m.array[m.parent(k)], m.array[k]
		k = m.parent(k)
	}
}

func (m *MaxHeap) ExtractMax() int {
	// 取出第一个元素，然后用最后一个元素代替第一个元素的值
	e := m.findMax()
	m.array[0] = m.array[len(m.array)-1]
	m.array = m.array[:len(m.array)-1]
	// 维护堆结构
	m.siftDown(0)
	return e
}

// 元素下沉
func (m *MaxHeap) siftDown(k int) {
	for m.leftChild(k) < len(m.array) { // 有子节点
		// 找出最大值的子节点
		j := m.leftChild(k)
		if j+1 < m.GetSize() && m.array[j] < m.array[j+1] {
			j++
		}
		// 当前节点就是最大值节点，则不需要进行任何操作
		if m.array[k] >= m.array[j] {
			break
		}
		// 当前父节点的值小于子节点，需要进行节点交换
		m.array[k], m.array[j] = m.array[j], m.array[k]
		k = j
	}
}

// 取出最大元素，同时添加一个元素
func (m *MaxHeap) Replace(e int) int {
	me := m.findMax()
	m.array[0] = e
	m.siftDown(0)
	return me
}

// 将任意数组整理成堆的形状
func (m *MaxHeap) Heapify(arr []int) {
	m.array = make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		m.array[i] = arr[i]
	}
	// 从最后一个元素的父节点向前开始做下沉操作；即最后一个有子节点的节点
	for k := m.parent(len(arr) - 1); k >= 0; k-- {
		m.siftDown(k)
	}
}
