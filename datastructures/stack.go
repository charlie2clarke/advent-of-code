package datastructures

type Stack[T any] []T

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Pop() T {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack[T]) Peek() T {
	return (*s)[len(*s)-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

func (s *Stack[T]) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *Stack[T]) Reverse() {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		s.Swap(i, j)
	}
}

func (s *Stack[T]) Copy() *Stack[T] {
	s2 := NewStack[T]()
	*s2 = append(*s2, *s...)
	return s2
}
