package datastructures

import "fmt"

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(values ...T) {
	for _, value := range values {
		s[value] = struct{}{}
	}
}

func (s Set[T]) Remove(value T) {
	delete(s, value)
}

func (s Set[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Clear() {
	for key := range s {
		delete(s, key)
	}
}

func (s Set[T]) Values() []T {
	values := make([]T, 0, len(s))
	for value := range s {
		values = append(values, value)
	}
	return values
}

func (s Set[T]) String() string {
	return fmt.Sprint(s.Values())
}

func (s Set[T]) IsEqual(other Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	for value := range s {
		if !other.Contains(value) {
			return false
		}
	}
	return true
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	union := NewSet[T]()
	for value := range s {
		union.Add(value)
	}
	for value := range other {
		union.Add(value)
	}
	return union
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	intersection := NewSet[T]()
	for value := range s {
		if other.Contains(value) {
			intersection.Add(value)
		}
	}
	return intersection
}

func (s Set[T]) Difference(other Set[T]) Set[T] {
	difference := NewSet[T]()
	for value := range s {
		if !other.Contains(value) {
			difference.Add(value)
		}
	}
	return difference
}

// Returns true if all items in the set are present in the other set.
func (s Set[T]) IsSubset(other Set[T]) bool {
	for value := range s {
		if !other.Contains(value) {
			return false
		}
	}
	return true
}

// Returns true of no items in the set are present in the other set.
func (s Set[T]) IsDisjoint(other Set[T]) bool {
	for value := range s {
		if other.Contains(value) {
			return false
		}
	}
	return true
}

func (s Set[T]) Copy() Set[T] {
	copy := NewSet[T]()
	for value := range s {
		copy.Add(value)
	}
	return copy
}

// Returns as set that contais all items from both sets except items that are present in both sets.
func (s Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	return s.Union(other).Difference(s.Intersection(other))
}

func (s Set[T]) AddAll(other Set[T]) {
	for value := range other {
		s.Add(value)
	}
}

func (s Set[T]) RemoveAll(other Set[T]) {
	for value := range other {
		s.Remove(value)
	}
}

func (s Set[T]) RetainAll(other Set[T]) {
	for value := range s {
		if !other.Contains(value) {
			s.Remove(value)
		}
	}
}
