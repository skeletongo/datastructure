// 一维线段树
package segment

type Segment interface {
	Set(index int, data interface{})
	Query(ql, qr int) interface{}
}
