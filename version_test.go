package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionIsCorrect(t *testing.T) {
	assert.Equal(t, version, "v0.0.1")
}

func TestNameIsCorrect(t *testing.T) {
	assert.Equal(t, name, "set")
}
