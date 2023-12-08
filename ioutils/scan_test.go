package ioutils

import (
	"bufio"
	"fmt"
	"os"
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

func BenchmarkScanner(b *testing.B) {
	file, err := os.Open("/var/log/system.log")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	b.Run("ReadLineFunc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReadLineFunc(file, func(num int, line string) bool {
				return true
			})
		}
	})

	b.Run("ReadLineBytesFunc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ReadLineBytesFunc(file, func(num int, line []byte) bool {
				return true
			})
		}
	})

	b.Run("bufioScanner_string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				scanner.Text()
			}
		}
	})

	b.Run("bufioScanner_bytes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				scanner.Bytes()
			}
		}
	})

}
