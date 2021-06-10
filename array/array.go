// 动态数组，其实go语言本身的slice就是动态数组,这里只是演示动态数组这种数据结构
package array

import (
	"bytes"
	"fmt"
)

// Array 动态数组
type Array struct {
	// 长度
	len int

	// 因为go语言的数组元素个数也是数组类型的一部分
	// 所以这里不能用数组声明，否则不能扩容
	// 这里为了演示用slice代替
	// 在实际开发中直接用slice就可以了，slice本身就是动态数组
	data []interface{}
}

// New 创建动态数组
// capacity 初始容量
func New(capacity ...int) *Array {
	_cap := 1
	if len(capacity) > 0 && capacity[0] > 0 {
		_cap = capacity[0]
	}
	return &Array{data: make([]interface{}, _cap)}
}

// Len 获取长度
func (a *Array) Len() int {
	return a.len
}

// Cap 获取容量
func (a *Array) Cap() int {
	return cap(a.data)
}

// IsEmpty 是否为空
func (a *Array) IsEmpty() bool {
	return a.len == 0
}

// Insert 在任意位置插入元素
func (a *Array) Insert(index int, v interface{}) {
	if index < 0 || index > a.len {
		panic("index out of bounds")
	}
	// 扩容
	if a.len == cap(a.data) {
		a.resize(2 * a.len)
	}
	// 插入元素,从插入位置开始所有元素后移一个位置
	for i := a.len - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = v
	a.len++
}

// resize 容量调整
func (a *Array) resize(n int) {
	data := make([]interface{}, n)
	for i := 0; i < a.len; i++ {
		data[i] = a.data[i]
	}
	a.data = data
}

// AddFirst 在第一个位置插入元素
func (a *Array) AddFirst(v interface{}) {
	a.Insert(0, v)
}

// AddLast 在最后一个位置添加元素
func (a *Array) AddLast(v interface{}) {
	a.Insert(a.len, v)
}

// Remove 移除指定位置的元素
func (a *Array) Remove(index int) interface{} {
	if index < 0 || index >= a.len {
		panic("index out of bounds")
	}
	e := a.data[index]
	// 移除
	for i := index + 1; i < a.len; i++ {
		a.data[i-1] = a.data[i]
	}
	a.len--
	a.data[a.len] = nil
	// 缩容，长度等于四分之一容量时，将容量缩小一半，但容量不能为零
	// 假如当长度是原来容量的一半时缩小一半容量，会有时间复杂度的震荡
	// 例如在刚刚扩容后移除一个元素会引起缩容，在刚刚移除一个元素后又添加一个元素又会引起扩容
	if l := cap(a.data); a.len == l/4 && l/2 != 0 {
		a.resize(l / 2)
	}
	return e
}

// RemoveFirst 移除第一个元素
func (a *Array) RemoveFirst() interface{} {
	return a.Remove(0)
}

// RemoveLast 移除最后一个元素
func (a *Array) RemoveLast() interface{} {
	return a.Remove(a.len - 1)
}

// Set 设置或修改某个位置的元素
func (a *Array) Set(index int, v interface{}) {
	if index < 0 || index >= a.len {
		panic("index out of bounds")
	}
	a.data[index] = v
}

// Swap 交换两个元素的位置
func (a *Array) Swap(i, j int) {
	if i < 0 || i >= a.Len() {
		return
	}
	if j < 0 || j >= a.Len() {
		return
	}
	if i == j {
		return
	}
	a.data[i], a.data[j] = a.data[j], a.data[i]
}

// ContainsByFunc 是否包含某个元素
func (a *Array) Contains(e interface{}, f func(a, b interface{}) bool) bool {
	for i := 0; i < a.len; i++ {
		if f(a.data[i], e) {
			return true
		}
	}
	return false
}

// Get 获取指定位置的元素
func (a *Array) Get(index int) interface{} {
	if index < 0 || index >= a.len {
		panic("index out of bounds")
	}
	return a.data[index]
}

func (a *Array) String() string {
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("len: %d cap: %d", a.len, cap(a.data)))
	buf.WriteString(" array: [")
	if a.len > 0 {
		buf.WriteString(fmt.Sprintf("%v", a.data[0]))
	}
	for i := 1; i < a.len; i++ {
		buf.WriteString(fmt.Sprintf(", %v", a.data[i]))
	}
	buf.WriteString("]")
	return buf.String()
}
