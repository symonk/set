package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElementsAreDistinct(t *testing.T) {
	a := New(5, 1, 2, 3, 4, 5)
	a.Add(2)
	a.Add(3)
	a.Add(3)
	assert.True(t, a.Len() == 5)

	for i := 0; i < a.Len(); i++ {
		random, _ := a.Pop()
		assert.True(t, !a.Contains(random))
	}
}

func TestInstantiationWithElements(t *testing.T) {
	set := New(5, 1, 2, 3, 4, 5)
	assert.Equal(t, set.Len(), 5)
}

func TestWithNoElements(t *testing.T) {
	set := New[string](0)
	assert.Zero(t, set.Len())
	assert.True(t, set.IsEmpty())
}

func TestClearingMapIsCorrect(t *testing.T) {
	set := New(3, "foo", "bar", "baz")
	assert.False(t, set.IsEmpty())
	set.Clear()
	assert.Zero(t, set.Len())
}

func TestRemoveAllElements(t *testing.T) {
	set := New(3, "foo", "bar", "baz")
	set.Remove("foo")
	assert.Equal(t, set.Len(), 2)
	assert.False(t, set.Contains("foo"))
	assert.True(t, set.Contains("bar"))
	assert.True(t, set.Contains("baz"))
}

func TestRemoveNonExistingIsNoOp(t *testing.T) {
	set := New[string](0)
	set.Remove("notinset")
	assert.Zero(t, set.Len())
}

func TestIsANewCopy(t *testing.T) {
	original := New(4, 1, 2, 3, 4)
	cp := original.Copy()
	original.Remove(4)
	assert.True(t, cp.Contains(1))
	assert.True(t, cp.Contains(2))
	assert.True(t, cp.Contains(3))
	assert.True(t, cp.Contains(4))

	assert.False(t, original.Contains(4))
}

func TestPopRemovesARandomElement(t *testing.T) {
	set := New(3, 1, 2, 3)
	// flush out the entire set
	for i := 0; i < 3; i++ {
		element, err := set.Pop()
		assert.NotZero(t, element)
		assert.Nil(t, err)
	}

	falsy, err := set.Pop()
	assert.Zero(t, falsy)
	assert.ErrorIs(t, err, ErrPopFromEmptySet)
	assert.ErrorContains(t, err, "pop from an empty set")
}

func TestIsNotDisjoint(t *testing.T) {
	a := New(5, 1, 2, 3, 4, 5)
	b := New(5, 5, 4, 3, 2, 1)
	isDisjoint := a.IsDisjoint(b)
	assert.False(t, isDisjoint)
}

func TestIsDisjoint(t *testing.T) {
	assert.True(t, New(3, 1, 2, 3).IsDisjoint(New(4, 5, 6)))
}

func TestIsSubsetYes(t *testing.T) {
	a := New(3, 1, 2, 3)
	b := New(3, 2, 3, 4)
	assert.False(t, a.IsSubset(b))
	assert.False(t, b.IsSubset(a))
}

func TestIsSubsetNo(t *testing.T) {
	a := New(5, 1, 2, 3, 4, 5)
	b := New(10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	assert.True(t, a.IsSubset(b))
	assert.False(t, b.IsSubset(a))
}

func TestEmptySetSubset(t *testing.T) {
	assert.True(t, New[int](0).IsSubset(New[int](0)))
}

func TestIsSuperSetYes(t *testing.T) {
	a := New(3, 1, 2, 3)
	b := New(3, 3, 2, 1)
	assert.True(t, a.IsSuperSet(b))
	assert.True(t, b.IsSuperSet(a))

	c := New(3, 1, 2, 3, 4, 5)
	d := New(3, 2, 3, 4)
	assert.True(t, c.IsSuperSet(d))
}

func TestIsSuperSetNo(t *testing.T) {
	a := New(3, 1, 2, 3)
	b := New(3, 2, 3)
	assert.False(t, b.IsSuperSet(a))
}

func TestEmptySetSuperset(t *testing.T) {
	assert.True(t, New[string](0).IsSuperSet(New[string](0)))
}

func TestUnionEmptySet(t *testing.T) {
	a := New[int](0)
	b := New[int](0)
	c := New[int](0)
	assert.True(t, a.Union(b).Equals(c))
}

func TestUnionMergesSuccessfully(t *testing.T) {
	a := New(5, 1, 2, 3, 4, 5)
	b := New(3, 1, 2, 3)
	c := New(8, 1, 2, 3, 4, 5, 6, 7, 8)
	union := a.Union(b, c)
	expected := New(3, 1, 2, 3, 4, 5, 6, 7, 8)
	assert.True(t, union.Equals(expected))
}

func TestUnionEmpty(t *testing.T) {
	a := New(3, 1, 2, 3)
	b := a.Union(New[int](0))
	assert.True(t, b.Equals(a))
}

func TestUpdateEmpty(t *testing.T) {
	a := New(3, 1, 2, 3)
	a.Update(New[int](0))
	assert.True(t, a.Equals(New(3, 1, 2, 3)))
}

func TestUpdateMultiple(t *testing.T) {
	a := New(3, 1, 2, 3)
	a.Update(New(3, 5, 6, 7))
	a.Update(New(3, 8, 9, 10))
	a.Update(New(3, 11, 12, 13))
	a.Update(New(3, 14, 15, 16))
	a.Update(New(3, 17, 18, 19))
	expected := New(18, 1, 2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	assert.True(t, a.Equals(expected))
}

func TestDifferentLensEqualsFalse(t *testing.T) {
	a := New(0, 1, 2, 3)
	b := New(0, 1, 2)
	assert.False(t, a.Equals((b)))
}
