package set

// Set is a generic implementation of a Hashset.
type Set[T comparable] struct {
	store map[T]struct{}
}

// Add adds the specified element into the set
func (s *Set[T]) Add(element T) {
	s.store[element] = struct{}{}
}
