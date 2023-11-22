package ioutils

import (
	"fmt"
	"strings"
	"testing"
)

func TestReadLineFunc(t *testing.T) {
	reader := strings.NewReader(`
1
2

3
`)

	err := ReadLineFunc(reader, func(num int, line string) bool {
		fmt.Printf("num: %d, line: %s\n", num, line)
		//if num == 2 {
		//	return false
		//}
		return true
	})
	t.Logf("err: %v", err)
}
