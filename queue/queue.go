package queue

type Queue[T any] interface {
	Push(v T) T
	Pop() (T, error)
	Empty() bool
	Head() (T, error)
}
