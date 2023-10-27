package datastructures

type Queue[T any] []T

func (q *Queue[T]) Enqueue(item T) {
	*q = append(*q, item)
}

func (q *Queue[T]) Dequeue() T {
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

func (q *Queue[T]) Peek() T {
	return (*q)[0]
}

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue[T]) Size() int {
	return len(*q)
}

func (q *Queue[T]) Clear() {
	*q = []T{}
}
