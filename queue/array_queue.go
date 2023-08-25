package queue

import "errors"

var (
	ErrQueueEmpty = errors.New("队列为空")
)

// ArrayQueue 泛型队列
type ArrayQueue[T any] struct {
	list []T
}

// NewArrayQueue 初始化泛型队列
func NewArrayQueue[T any]() *ArrayQueue[T] {
	list := make([]T, 0)
	return &ArrayQueue[T]{
		list: list,
	}
}

// Push 向队列中添加数据
func (queue *ArrayQueue[T]) Push(v T) T {
	queue.list = append(queue.list, v)
	return v
}

// Pop 获取队列中第一条数据，并将其删除
// 如何这个队列为空时会返回零值及异常
func (queue *ArrayQueue[T]) Pop() (T, error) {
	if len(queue.list) < 1 {
		var zero T
		return zero, ErrQueueEmpty
	}
	v := queue.list[:1][0]
	queue.list = queue.list[1:len(queue.list)]
	return v, nil
}

// Empty 队列是否为空
func (queue *ArrayQueue[T]) Empty() bool {
	return len(queue.list) == 0
}

// Head 获取队列头部数据
func (queue *ArrayQueue[T]) Head() (T, error) {
	if queue.Empty() {
		var zero T
		return zero, ErrQueueEmpty
	}
	return queue.list[:1][0], nil
}
