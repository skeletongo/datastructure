package queue

// Queue 队列
type Queue interface {
	Enqueue(interface{})
	Dequeue() interface{}
}
