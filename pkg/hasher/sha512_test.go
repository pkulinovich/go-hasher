package hasher

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewSHA512(t *testing.T) {
	assert.IsType(t, &SHA512{}, NewSha512())
}

func TestSHA512_Hash(t *testing.T) {
	hasher := NewSha512()

	r := strings.NewReader("abcde")

	hash, err := hasher.Hash(r)
	if assert.Nil(t, err) {
		assert.Equal(t, "878ae65a92e86cac011a570d4c30a7eaec442b85ce8eca0c2952b5e3cc0628c2e79d889ad4d5c7c626986d452dd86374b6ffaa7cd8b67665bef2289a5c70b0a1", hash)
	}
}
