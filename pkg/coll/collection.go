package coll

type Collection[T any] interface {
	Len() int
	Contains(v T) bool
	Append(T)
	Remove(T)
}
