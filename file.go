package goutils

import (
	"bufio"
	"os"
)

// ReadLineFunc read the file line by line and call f(c) to process each line of data
func ReadLineFunc(path string, f func(line string) error) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err = f(BytesToString(scanner.Bytes())); err != nil {
			return err
		}
	}
	return err
}

// IsDir reports whether this path is a directory
func IsDir(path string) bool {
	if f, err := os.Stat(path); err != nil {
		return false
	} else {
		return f.IsDir()
	}
}
