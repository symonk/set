package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionIsCorrect(t *testing.T) {
	assert.Equal(t, Version, "v0.0.1")
}
