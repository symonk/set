package set

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
// if
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
