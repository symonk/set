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
