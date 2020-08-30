package queue

// Queue 队列
type Queue interface {
	Len() int
	Enqueue(interface{})
	Dequeue() interface{}
}
