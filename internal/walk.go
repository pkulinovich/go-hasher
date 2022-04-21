package internal

import (
	"os"
	"path/filepath"
)

func Walk(inputPath string) chan ObjectInfo {
	resCh := make(chan ObjectInfo, 1)
	go func(resCh chan ObjectInfo) {
		defer close(resCh)

		err := filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
			resCh <- ObjectInfo{
				Path: path,
				Info: info,
				Err:  err,
			}
			return nil
		})

		if err != nil {
			resCh <- ObjectInfo{
				Path: inputPath,
				Info: nil,
				Err:  err,
			}
		}
	}(resCh)
	return resCh
}
