package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstantiationWithElements(t *testing.T) {
	set := New(1, 2, 3, 4, 5)
	assert.Equal(t, set.Len(), 5)
}

func TestWithNoElements(t *testing.T) {
	set := New[string]()
	assert.Zero(t, set.Len())
}

func TestClearingMapIsCorrect(t *testing.T) {
	set := New("foo", "bar", "baz")
	set.Clear()
	assert.Zero(t, set.Len())
}

func TestRemoveAllElements(t *testing.T) {
	set := New("foo", "bar", "baz")
	set.Remove("foo")
	assert.Equal(t, set.Len(), 2)
	assert.False(t, set.Contains("foo"))
	assert.True(t, set.Contains("bar"))
	assert.True(t, set.Contains("baz"))
}

func TestRemoveNonExistingIsNoOp(t *testing.T) {
	set := New[string]()
	set.Remove("notinset")
	assert.Zero(t, set.Len())
}

func TestIsANewCopy(t *testing.T) {
	original := New(1, 2, 3, 4)
	cp := original.Copy()
	original.Remove(4)
	assert.True(t, cp.Contains(1))
	assert.True(t, cp.Contains(2))
	assert.True(t, cp.Contains(3))
	assert.True(t, cp.Contains(4))

	assert.False(t, original.Contains(4))
}
