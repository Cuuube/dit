package coll

type empty struct{}

var _ Collection[string] = (*Set[string])(nil)

type Set[T DictKeyable] map[T]struct{}

func (set *Set[T]) Len() int {
	s := *set
	return len(s)
}

func (set *Set[T]) Append(v T) {
	s := *set
	s[v] = empty{}
}

func (set *Set[T]) Contains(v T) bool {
	s := *set
	_, found := s[v]
	return found
}

func (set *Set[T]) Remove(v T) {
	s := *set
	delete(s, v)
}
