package list

// LinkedList 单向链表
type LinkedList[T any] struct {
	// 链表的头节点
	head *node[T]
	size int
}

// node 链表节点
type node[T any] struct {
	value T
	next  *node[T]
}

// NewLinkedList 初始化一个单项链表
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
		size: 0,
	}
}

func (l *LinkedList[T]) Add(v T) {
	n := l.head
	if n == nil {
		var nextNode node[T]
		nextNode.value = v
		l.head = &nextNode
		l.size = l.size + 1
		return
	}
	// 循环遍历链表下一节点信息，如果下一节点为空时赋值
	for {
		pre := n
		n := n.next
		if n == nil {
			var nextNode node[T]
			nextNode.value = v
			pre.next = &nextNode
			l.size = l.size + 1
			break
		}

	}

}

func (l *LinkedList[T]) AddAll(v ...T) {
	panic("implement me")
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList[T]) FindFist(v T) int {
	panic("implement me")
}

func (l *LinkedList[T]) FindAll(v T) []int {
	panic("implement me")
}

// Get 根据坐标获取值
// 当坐标值大于链表的大小时抛出 ErrArrayIndexOut 数组越界异常
// 当链表数据为空时，抛出 ErrListIsEmpty 数组为空异常
func (l *LinkedList[T]) Get(i int) (T, error) {
	var zero T
	// 当坐标大于链表大小时抛出异常
	if l.size < i {
		return zero, ErrArrayIndexOut
	}
	// 当链表中数据为空时抛出异常
	if l.size == 0 {
		return zero, ErrListIsEmpty
	}
	n := l.head
	if i == 0 {
		return n.value, nil
	}
	inx := 1
	for {
		n := n.next
		if inx == i {
			return n.value, nil
		}
	}
}
