package queue

import "errors"

var (
	ErrQueueEmpty = errors.New("队列为空")
)

// Queue 泛型队列
type Queue[T any] struct {
	list []T
}

// NewQueue 初始话泛型队列
func NewQueue[T any]() *Queue[T] {
	list := make([]T, 0)
	return &Queue[T]{
		list: list,
	}
}

// Push 向队列中添加数据
func (queue *Queue[T]) Push(v T) T {
	queue.list = append(queue.list, v)
	return v
}

// Pop 获取队列中第一条数据
// 如何这个队列为空时会返回零值及异常
func (queue *Queue[T]) Pop() (T, error) {
	if len(queue.list) < 1 {
		var zero T
		return zero, ErrQueueEmpty
	}
	v := queue.list[:1][0]
	queue.list = queue.list[1:len(queue.list)]
	return v, nil
}
