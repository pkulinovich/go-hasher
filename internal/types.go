package internal

import "os"

type ObjectInfo struct {
	Path string
	Info os.FileInfo
	Err  error
}

type HashMaker interface {
	HashFile(path string) (string, error)
}
