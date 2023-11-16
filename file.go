package goutils

import (
	"os"
)

// ReadFileLineFunc read the file line by line and call f(c) to process each line of data
func ReadFileLineFunc(path string, f func(num int, line string) bool) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return ReadLineFunc(file, f)
}

// IsDir reports whether this path is a directory
func IsDir(path string) bool {
	if f, err := os.Stat(path); err != nil {
		return false
	} else {
		return f.IsDir()
	}
}
