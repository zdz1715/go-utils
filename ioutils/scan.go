package ioutils

import (
	"bufio"
	goutils "github.com/zdz1715/go-utils/goutils"
	"io"
)

// ReadLineFunc read the io.Reader line by line and call f(c) to process each line of string
func ReadLineFunc(reader io.Reader, f func(num int, line string) bool) error {
	scanner := bufio.NewScanner(reader)
	num := 0
	for scanner.Scan() {
		num++
		if !f(num, goutils.BytesToString(scanner.Bytes())) {
			break
		}
	}
	return scanner.Err()
}

// ReadLineBytesFunc read the io.Reader line by line and call f(c) to process each line of bytes
func ReadLineBytesFunc(reader io.Reader, f func(num int, line []byte) bool) error {
	scanner := bufio.NewScanner(reader)
	num := 0
	for scanner.Scan() {
		num++
		if !f(num, scanner.Bytes()) {
			break
		}
	}
	return scanner.Err()
}
