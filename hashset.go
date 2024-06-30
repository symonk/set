package set

import (
	"errors"
	"reflect"
)

/*
Difference()
DifferenceUpdate()
Intersection()
IntersectionUpdate()
SymmetricDifference()
SymmetricDifferenceUpdate()
Update()
*/

type nothing struct{}

var sentinelNothing = nothing{}

// ErrPopFromEmptySet is returned when calling Pop() on an empty set
var ErrPopFromEmptySet = errors.New("pop from an empty set")

// Set is a generic implementation of a Hashset.
type Set[T comparable] struct {
	elements map[T]nothing
}

// New returns a new generic hashset of type T.
func New[T comparable](capacity int, elements ...T) *Set[T] {
	s := &Set[T]{
		elements: make(map[T]nothing, capacity),
	}
	for _, element := range elements {
		s.Add(element)
	}
	return s

}

// Add adds the specified element into this set.
// Add has no effect if the element is already present.
func (s *Set[T]) Add(element T) {
	s.elements[element] = sentinelNothing
}

// Remove deletes the element from this set.
// if the element is not in the set it is
// a no op
func (s *Set[T]) Remove(element T) {
	delete(s.elements, element)
}

// Clear removes all elements from this set.
func (s *Set[T]) Clear() {
	clear(s.elements)
}

// Len returns the number of elements in this hashset.
func (s *Set[T]) Len() int {
	return len(s.elements)
}

// IsEmpty returns true if this set contains no elements.
func (s *Set[T]) IsEmpty() bool {
	return s.Len() == 0
}

// Contains returns true if the element is in this set.
func (s *Set[T]) Contains(element T) bool {
	_, existed := s.elements[element]
	return existed
}

// Copy returns a copy of this set
func (s *Set[T]) Copy() *Set[T] {
	newCopy := make(map[T]nothing, len(s.elements))
	for k, v := range s.elements {
		newCopy[k] = v
	}
	return &Set[T]{
		elements: newCopy,
	}
}

// Pop returns an arbitrary element from this set
// if this set is empty, an error is returned
func (s *Set[T]) Pop() (T, error) {
	for element := range s.elements {
		s.Remove(element)
		return element, nil
	}
	var falsy T
	return falsy, ErrPopFromEmptySet
}

// IsDisjoint returns true if this set and other
// have a null intersection.
func (s *Set[T]) IsDisjoint(other *Set[T]) bool {
	for element := range other.elements {
		if s.Contains(element) {
			return false
		}
	}
	return true
}

// IsSubset tests whether every element in this set
// is contained in the other set.
// empty sets are considered subsets of each other.
func (s *Set[T]) IsSubset(other *Set[T]) bool {
	if other.Len() < s.Len() {
		return false
	}

	for element := range s.elements {
		if !other.Contains(element) {
			return false
		}
	}
	return true

}

// IsSuperset tests whether every element in other is
// contained within this set.
// empty sets are considered supersets of each other.
func (s *Set[T]) IsSuperSet(other *Set[T]) bool {
	return other.IsSubset(s)
}

// Equals returns true if both sets have the exact same
// elements
func (s *Set[T]) Equals(other *Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	return reflect.DeepEqual(s.elements, other.elements)
}

// Union returns a new set containing all of the elements
// that appear in all of the sets.
// empty sets are not included in the union.
func (s *Set[T]) Union(others ...*Set[T]) *Set[T] {
	result := New[T](0)
	counter := make(map[T]int)

	var target int
	if s.Len() != 0 {
		target = 1
	} else {
		target = 0
	}

	for element := range s.elements {
		counter[element] += 1
	}
	for _, set := range others {
		if set.Len() != 0 {
			target++
		}
		for element := range set.elements {
			counter[element] += 1
		}
	}
	for element, count := range counter {
		if count == target {
			result.Add(element)
		}
	}

	return result
}
