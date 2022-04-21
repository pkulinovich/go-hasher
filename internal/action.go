package internal

import (
	"errors"
	"os"

	"go-hasher/pkg/hasher"
)

type Action struct {
	hasher hasher.Hasher
}

func NewAction(algo string) (*Action, error) {
	h, err := hasher.New(algo)
	if err != nil {
		return nil, err
	}
	return &Action{hasher: h}, nil
}

func (s *Action) HashFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", errors.New("wrong file")
	}
	defer file.Close()

	res, err := s.hasher.Hash(file)
	if err != nil {
		return "", err
	}
	return res, nil
}
