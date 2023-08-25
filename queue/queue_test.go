package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	MockValues       = []int64{1, 2, 3, 5}
	FirstValue int64 = 1
)

func TestNewQueue(t *testing.T) {
	queue := NewArrayQueue[int]()
	queue.Push(1)
	t.Logf("queue push data after,queue list :%+v", queue.list)
	v, err := queue.Pop()
	if err != nil {
		t.Errorf("队列获取数据异常，异常信息为:%v", err)
		return
	}
	t.Logf("queue pop data after,queue list :%+v", queue.list)
	assert.Equal(t, v, 1, "队列中获取的数据与导入的数据不相符")
}

func TestArrayQueue_Head(t *testing.T) {
	queue := NewArrayQueue[int64]()
	for _, value := range MockValues {
		queue.Push(value)
	}
	head, err := queue.Head()
	if err != nil {
		t.Errorf("测试 Head 方法异常，异常信息为%+v", err)
		return
	}
	assert.Equal(t, FirstValue, head, "获取头部数据异常")
}
