package collection

import (
	"sync"
	"time"
)

// RollingWindow 滑动窗口
type RollingWindow struct {
	sync.RWMutex
	size     int                      // 窗口总数
	offset   int                      // 窗口索引
	lastTime time.Time                // 最后使用的窗口的开始时间
	interval time.Duration            // 窗口时长
	create   func() interface{}       // 窗口对象创建方法
	reset    func(bucket interface{}) // 窗口对象重置方法
	buckets  []interface{}            // 窗口对象
}

// NewRollingWindow 创建滑动窗口
// size 窗口总数
// interval 窗口时间间隔
// create 窗口对象创建方法，返回值必须是指针类型
// reset 窗口对象重置方法，方法参数为指针类型
func NewRollingWindow(size int, interval time.Duration,
	create func() interface{}, reset func(bucket interface{})) *RollingWindow {
	if size < 1 {
		panic("size must be greater than 0")
	}
	if interval < 1 {
		panic("interval must be greater than 0")
	}
	if create == nil {
		panic("create not allow nil")
	}
	if reset == nil {
		panic("reset not allow nil")
	}

	ret := &RollingWindow{
		size:     size,
		lastTime: time.Now(),
		interval: interval,
		create:   create,
		reset:    reset,
	}
	for i := 0; i < size; i++ {
		ret.buckets = append(ret.buckets, create())
	}
	return ret
}

// 从上次更新窗口对象到当前时间已经多少个窗口过期了
func (r *RollingWindow) timeoutBucketNumber(t time.Time) int {
	n := int(t.Sub(r.lastTime) / r.interval)
	if 0 <= n && n < r.size {
		return n
	}
	return r.size
}

// Range 遍历窗口对象
// f 对象处理方法，参数为指针类型
// ignoreCurrent 是否忽略当前窗口对象
// 线程安全
func (r *RollingWindow) Range(f func(bucket interface{}), ignoreCurrent ...bool) {
	r.RLock()
	defer r.RUnlock()

	var ignore bool
	if len(ignoreCurrent) > 0 {
		ignore = ignoreCurrent[0]
	}

	var count int
	n := r.timeoutBucketNumber(time.Now())
	if n == 0 && ignore {
		count = r.size - 1
	} else {
		count = r.size - n
	}

	for i := 0; i < count; i++ {
		f(r.buckets[(r.offset+n+i+1)%r.size])
	}
}

// Update 更新当前窗口对象
// f 当前窗口对象处理方法，参数为指针类型
// 线程安全
func (r *RollingWindow) Update(f func(bucket interface{})) {
	r.Lock()
	defer r.Unlock()

	// 更新窗口
	now := time.Now()
	n := r.timeoutBucketNumber(now)
	if n > 0 {
		// 更新窗口起始时间
		r.lastTime = now.Add(-now.Sub(r.lastTime) % r.interval)
		// 重置过期的窗口对象，否则遍历窗口对象时应该过期的对象数据还是存在的
		for i := 0; i < n; i++ {
			r.reset(r.buckets[(r.offset+i+1)%r.size])
		}
		// 更新窗口索引
		r.offset = (r.offset + n) % r.size
	}
	// 更新数据
	f(r.buckets[r.offset])
}
