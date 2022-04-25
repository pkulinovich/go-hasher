package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUnknown(t *testing.T) {
	hasher, err := New("something")
	if assert.Nil(t, hasher) && assert.Error(t, err) {
		assert.Equal(t, "unknown algorithm", err.Error())
	}
}

func TestNew_MD5(t *testing.T) {
	h, err := New("md5")
	if assert.Nil(t, err) {
		assert.IsType(t, &MD5{}, h)
	}
}

func TestNew_Sha512(t *testing.T) {
	h, err := New("sha512")
	if assert.Nil(t, err) {
		assert.IsType(t, &SHA512{}, h)
	}
}
