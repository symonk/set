package set

import "errors"

/*
	Difference()
	DifferenceUpdate()
	Intersection()
	IntersectionUpdate()
	IsDisjoint()
	IsSubset()
	IsSuperset()
	Pop()
	SymmetricDifference()
	SymmetricDifferenceUpdate()
	Union()
	Update()
*/

var ErrPopFromEmptySet = errors.New("pop from an empty set")

// Set is a generic implementation of a Hashset.
type Set[T comparable] struct {
	store map[T]struct{}
}

// New returns a new generic hashset of type T.
func New[T comparable](elements ...T) *Set[T] {
	s := &Set[T]{
		store: make(map[T]struct{}),
	}
	for _, element := range elements {
		s.Add(element)
	}
	return s

}

// Add adds the specified element into this set.
// Add has no effect if the element is already present.
func (s *Set[T]) Add(element T) {
	s.store[element] = struct{}{}
}

// Remove deletes the element from this set.
// if the element is not in the set it is
// a no op
func (s *Set[T]) Remove(element T) {
	delete(s.store, element)
}

// Clear removes all elements from this set.
func (s *Set[T]) Clear() {
	clear(s.store)
}

// Len returns the number of elements in this hashset.
func (s *Set[T]) Len() int {
	return len(s.store)
}

// Contains returns true if the element is in this set.
func (s *Set[T]) Contains(element T) bool {
	_, ok := s.store[element]
	return ok
}

// Copy returns a copy of this set
func (s *Set[T]) Copy() *Set[T] {
	newCopy := make(map[T]struct{}, len(s.store))
	for k, v := range s.store {
		newCopy[k] = v
	}
	return &Set[T]{
		store: newCopy,
	}
}

// Pop returns an arbitrary element from this set
// if this set is empty, an error is returned
func (s *Set[T]) Pop() (T, error) {
	for element := range s.store {
		s.Remove(element)
		return element, nil
	}
	var falsy T
	return falsy, ErrPopFromEmptySet
}
