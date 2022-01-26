// Package segment 一维线段树
package segment

type Segment interface {
	// Set 根据索引修改原数据
	Set(index int, data interface{})
	// Query 查询指定区间内的统计信息
	Query(ql, qr int) interface{}
}
