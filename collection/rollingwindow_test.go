package collection

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type Bucket struct {
	N float64
}

func create() interface{} {
	return &Bucket{}
}

func reset(bucket interface{}) {
	bucket.(*Bucket).N = 0
}

func update(bucket interface{}) {
	bucket.(*Bucket).N++
}

func TestNewRollingWindow(t *testing.T) {
	assert.NotNil(t, NewRollingWindow(10, time.Second, create, reset))

	assert.Panics(t, func() {
		NewRollingWindow(0, time.Second, create, reset)
	})

	assert.Panics(t, func() {
		NewRollingWindow(10, 0, create, reset)
	})

	assert.Panics(t, func() {
		NewRollingWindow(10, time.Second, nil, reset)
	})

	assert.Panics(t, func() {
		NewRollingWindow(10, time.Second, create, nil)
	})
}

func TestRollingWindow_Update(t *testing.T) {
	var r *RollingWindow
	buckets := make([]float64, 0, 3)
	f := func() []float64 {
		buckets = buckets[:0]
		r.Range(func(bucket interface{}) {
			buckets = append(buckets, bucket.(*Bucket).N)
		})
		return buckets
	}

	interval := time.Millisecond * 100
	r = NewRollingWindow(3, interval, create, reset)

	assert.Equal(t, []float64{0, 0, 0}, f())
	r.Update(update)
	assert.Equal(t, []float64{0, 0, 1}, f())

	time.Sleep(interval)
	r.Update(update)
	r.Update(update)
	assert.Equal(t, []float64{0, 1, 2}, f())

	time.Sleep(interval)
	r.Update(update)
	r.Update(update)
	r.Update(update)
	assert.Equal(t, []float64{1, 2, 3}, f())

	time.Sleep(interval)
	r.Update(update)
	assert.Equal(t, []float64{2, 3, 1}, f())
}

func TestRollingWindow_Update2(t *testing.T) {
	var r *RollingWindow
	buckets := make([]float64, 0, 3)
	f := func() []float64 {
		buckets = buckets[:0]
		r.Range(func(bucket interface{}) {
			buckets = append(buckets, bucket.(*Bucket).N)
		})
		return buckets
	}
	add := func(n float64) func(bucket interface{}) {
		return func(bucket interface{}) {
			bucket.(*Bucket).N += n
		}
	}

	interval := time.Millisecond * 100
	r = NewRollingWindow(3, interval, create, reset)

	assert.Equal(t, []float64{0, 0, 0}, f())
	r.Update(add(1))
	assert.Equal(t, []float64{0, 0, 1}, f())

	time.Sleep(time.Millisecond * 150)
	r.Update(add(2))
	r.Update(add(3))
	assert.Equal(t, []float64{0, 1, 5}, f())

	time.Sleep(time.Millisecond * 80)
	r.Update(add(4))
	r.Update(add(5))
	r.Update(add(6))
	assert.Equal(t, []float64{1, 5, 15}, f())

	time.Sleep(time.Millisecond * 300)
	r.Update(add(7))
	r.Update(add(8))
	r.Update(add(9))
	assert.Equal(t, []float64{0, 0, 24}, f())
}

func TestRollingWindow_Reset(t *testing.T) {
	interval := time.Millisecond * 100
	r := NewRollingWindow(3, interval, create, reset)

	buckets := make([]float64, 0, 3)
	f := func() []float64 {
		buckets = buckets[:0]
		r.Range(func(bucket interface{}) {
			buckets = append(buckets, bucket.(*Bucket).N)
		})
		return buckets
	}

	r.Update(update)
	time.Sleep(interval)
	assert.Equal(t, []float64{0, 1}, f())

	time.Sleep(interval)
	assert.Equal(t, []float64{1}, f())

	time.Sleep(interval)
	assert.Equal(t, []float64{}, f())

	r.Update(update)
	time.Sleep(interval * 4)
	assert.Equal(t, []float64{}, f())
}

func TestRollingWindow_Range(t *testing.T) {
	const size = 4
	const interval = time.Millisecond * 100
	tests := []struct {
		id     int
		win    *RollingWindow
		ignore bool
		expect float64
	}{
		{
			id:     0,
			win:    NewRollingWindow(size, interval, create, reset),
			ignore: false,
			expect: 10,
		},
		{
			id:     1,
			win:    NewRollingWindow(size, interval, create, reset),
			ignore: true,
			expect: 4,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint("test", test.id), func(t *testing.T) {
			r := test.win
			for x := 0; x < size; x++ { // 0 1 2 3
				for i := 0; i <= x; i++ {
					r.Update(func(bucket interface{}) {
						bucket.(*Bucket).N += float64(i)
					})
				}
				if x < size-1 { // 0 1 2
					time.Sleep(interval)
				}
			}
			var result float64
			r.Range(func(bucket interface{}) {
				result += bucket.(*Bucket).N
			}, test.ignore)
			assert.Equal(t, test.expect, result)
		})
	}
}

func TestRollingWindow_Race(t *testing.T) {
	r := NewRollingWindow(3, time.Millisecond*100, create, reset)
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				r.Update(func(bucket interface{}) {
					bucket.(*Bucket).N += float64(rand.Int63())
				})
				time.Sleep(time.Millisecond * 50)
			}
		}
	}()
	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				r.Range(func(bucket interface{}) {})
			}
		}
	}()
	time.Sleep(time.Millisecond * 500)
	close(stop)
}
