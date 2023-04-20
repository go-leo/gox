package filex

import (
	"os"
)

// IsDir reports whether the named file is a directory.
func IsDir(filepath string) bool {
	f, err := os.Stat(filepath)
	if err != nil {
		return false
	}
	return f.IsDir()
}

// IsDirectory reports whether the named file is a directory.
var IsDirectory = IsDir

// IsExist returns a boolean indicating whether a file or directory exist.
func IsExist(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

// IsNotExist returns a boolean indicating whether a file or directory not exist.
func IsNotExist(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return false
	}
	return os.IsNotExist(err)
}
