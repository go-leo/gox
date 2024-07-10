package filepathx

import (
	"io/fs"
	"path/filepath"
)

// List is used to list the paths of all regular files under the specified root directory.
func List(root string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	return files, err
}
