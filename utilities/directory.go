package utilities

import (
	"os"
	"path/filepath"
)

//GetDirContents ...
func GetDirContents(directory string) []os.FileInfo {
	files := []os.FileInfo{}

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, info)
		return nil
	})

	Check(err)

	return files
}
