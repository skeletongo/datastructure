package array

import (
	"bytes"
	"fmt"
	"reflect"
)

// 动态数组默认初始容量
const DefaultCapacity = 10

// 动态数组
// golang中的切片就是动态数组,这里当成静态数组使用了
type Array struct {
	size int // 长度
	data []interface{}
}

// 构造方法
func NewArray(c ...int) *Array {
	capacity := DefaultCapacity
	if len(c) > 0 && c[0] > 0 {
		capacity = c[0]
	}
	arr := Array{size: 0}
	arr.data = make([]interface{}, capacity)
	return &arr
}

// 获取大小
func (a *Array) GetSize() int {
	return a.size
}

// 获取容量
func (a *Array) GetCapacity() int {
	return len(a.data)
}

// 是否为空
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

// 在第一个位置插入元素
func (a *Array) AddFirst(e interface{}) {
	a.Insert(0, e)
}

// 在最后一个位置添加元素
func (a *Array) AddLast(e interface{}) {
	a.Insert(a.size, e)
}

// 容量调整
func (a *Array) resize(n int) {
	data := make([]interface{}, n)
	for i := 0; i < a.size; i++ {
		data[i] = a.data[i]
	}
	a.data = data
}

// 在任意位置插入元素
func (a *Array) Insert(index int, e interface{}) {
	if index < 0 || index > a.size { // size 是最后一个可添加元素的位置，是可以取到的；如果大于size就会隔一个空位添加元素，这是不允许的；
		panic("out of bounds")
	}
	// 扩容
	if a.size == len(a.data) {
		a.resize(2 * a.size)
	}
	// 插入元素
	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = e
	a.size++
}

// 移除第一个元素
func (a *Array) RemoveFirst() interface{} {
	return a.Remove(0)
}

// 移除最后一个元素
func (a *Array) RemoveLast() interface{} {
	return a.Remove(a.size - 1)
}

// 移除指定位置的元素
func (a *Array) Remove(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("out of bounds")
	}
	res := a.data[index]
	// 移除
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}
	// 这里先减1 再置空最后一个元素，可以省掉一次减1运算
	a.size--
	a.data[a.size] = nil
	// 缩容，长度等于四分之一容量时，将容量缩小一半
	// 当且仅当 l = 1 时，l/2 等于 0 ，但容量不能为0
	if l := len(a.data); a.size == l/4 && l/2 != 0 {
		a.resize(l / 2)
	}
	return res
}

// 设置或修改莫个位置的元素
func (a *Array) Set(index int, e interface{}) {
	if index < 0 || index >= a.size {
		panic("out of bounds")
	}
	a.data[index] = e
}

// 交换两个元素的位置
func (a *Array) Swap(index1, index2 int) {
	if index1 >= 0 && index2 >= 0 && index1 < a.GetSize() && index2 < a.GetSize() && index1 != index2 {
		a.data[index1], a.data[index2] = a.data[index2], a.data[index1]
	}
}

// 是否包含莫个元素
func (a *Array) Contains(e interface{}) bool {
	for i := 0; i < a.size; i++ {
		if reflect.DeepEqual(a.data[i], e) {
			return true
		}
	}
	return false
}

func (a *Array) Get(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("out of bounds")
	}
	return a.data[index]
}

func (a *Array) String() string {
	buf := bytes.Buffer{}
	buf.WriteString("Array: [")
	for i := 0; i < a.size; i++ {
		buf.WriteString(fmt.Sprintf("%v", a.data[i]))
		if i < a.size-1 {
			buf.WriteString(",")
		}
	}
	buf.WriteString("]")
	buf.WriteString(fmt.Sprintf(" len:%d cap:%d", a.size, len(a.data)))
	return buf.String()
}
