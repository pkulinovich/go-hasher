package hasher

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMD5(t *testing.T) {
	assert.IsType(t, &MD5{}, NewMD5())
}

func TestMD5_Hash(t *testing.T) {
	hasher := NewMD5()

	r := strings.NewReader("abcde")

	hash, err := hasher.Hash(r)
	if assert.Nil(t, err) {
		assert.Equal(t, "ab56b4d92b40713acc5af89985d4b786", hash)
	}
}
