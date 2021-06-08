// 堆的实现也可参考标准库container中heap的实现；实际开发中请使用标准库
package heap

type Heap interface {
	Add(e interface{})
	ExtractMax() interface{}
}
