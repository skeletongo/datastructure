package queue

// IQueue 队列
type IQueue interface {
	Len() int
	Cap() int
	Enqueue(interface{})
	Dequeue() interface{}
}
