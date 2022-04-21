package hasher

import (
	"crypto/sha512"
	"encoding/hex"
	"io"
)

type SHA512 struct{}

func NewSha512() *SHA512 {
	return &SHA512{}
}

func (m *SHA512) Hash(file io.Reader) (string, error) {
	hash := sha512.New()

	_, err := io.Copy(hash, file)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
