package list

import "errors"

var (
	ErrArrayIndexOut = errors.New("数组越界异常")
	ErrListIsEmpty   = errors.New("数组为空异常")
)

type List[T any] interface {
	Add(v T)
	AddAll(v ...T)
	Size() int
	IsEmpty() bool
	FindFist(v T) int
	FindAll(v T) []int
	Get(i int) (T, error)
}
