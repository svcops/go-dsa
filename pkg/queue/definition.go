package queue

// Queue
// @Description: 队列定义
type Queue interface {
	Size() int

	Offer(interface{})

	Poll() interface{}
}
