package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewQueue(t *testing.T) {
	queue := NewQueue[int]()
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
