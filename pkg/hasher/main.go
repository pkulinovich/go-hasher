package hasher

import (
	"io"

	"github.com/pkg/errors"
)

type Hasher interface {
	Hash(file io.Reader) (string, error)
}

func New(algo string) (hasher Hasher, err error) {
	switch algo {
	case "md5", "MD5":
		hasher = NewMD5()
	case "sha512", "SHA512":
		hasher = NewSha512()
	default:
		err = errors.New("unknown algorithm")
	}
	return
}
