package fileutils

import (
	"errors"
	"github.com/zdz1715/go-utils/ioutils"
	"io"
	"os"
	"path/filepath"
)

// CreateIfNotExists creates a file or a directory only if it does not already exist.
func CreateIfNotExists(path string, isDir bool) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			if isDir {
				return os.MkdirAll(path, 0o755)
			}
			if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
				return err
			}
			f, err := os.OpenFile(path, os.O_CREATE, 0o755)
			if err != nil {
				return err
			}
			defer f.Close()
		}
	}
	return nil
}

// CopyFile copies from src to dst until either EOF is reached
// on src or an error occurs. It verifies src exists and removes
// the dst if it exists.
func CopyFile(src, dst string) (int64, error) {
	cleanSrc := filepath.Clean(src)
	cleanDst := filepath.Clean(dst)
	if cleanSrc == cleanDst {
		return 0, nil
	}
	sf, err := os.Open(cleanSrc)
	if err != nil {
		return 0, err
	}
	defer sf.Close()
	if err := os.Remove(cleanDst); err != nil && !os.IsNotExist(err) {
		return 0, err
	}
	df, err := os.Create(cleanDst)
	if err != nil {
		return 0, err
	}
	defer df.Close()
	return io.Copy(df, sf)
}

// ReadLineFunc read the file line by line and call f(c) to process each line of data
func ReadLineFunc(path string, f func(num int, line string) bool) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()
	return ioutils.ReadLineFunc(file, f)
}

// IsDir reports whether this path is a directory
func IsDir(path string) bool {
	return IsDirE(path) == nil
}

// IsDirE returns an error as to whether this path is a directory
func IsDirE(path string) error {
	f, err := os.Stat(path)
	if err != nil {
		return err
	}
	if f.IsDir() {
		return nil
	}
	return &os.PathError{Op: "IsDirE", Path: path, Err: errors.New("no such directory")}
}

// IsNotExist reports whether this path is not exist
func IsNotExist(path string) bool {
	b, _ := IsNotExistE(path)
	return b
}

// IsNotExistE returns whether this path is not exist and error
func IsNotExistE(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		return os.IsNotExist(err), err
	} else {
		return false, nil
	}
}
