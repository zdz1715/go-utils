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

func TestForwardScanner_Scan(t *testing.T) {
	//	str := `
	//789
	//456
	//123`
	//	reader := strings.NewReader(str)
	//size := int64(len(str))
	reader, err := os.Open("/var/log/system.log")
	if err != nil {
		t.Fatal(err)
	}
	fileInfo, err := reader.Stat()
	if err != nil {
		t.Fatal(err)
	}
	size := fileInfo.Size()
	scanner := NewForwardScanner(reader, size)

	for scanner.Scan() {
		fmt.Println("---", fmt.Sprintf("%q", scanner.Text()))
	}

	t.Logf("err: %v", scanner.Err())
}

func TestScanner_Scan(t *testing.T) {
	str := `
789
456
123
`
	reader := strings.NewReader(str)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println("---", fmt.Sprintf("%q", scanner.Text()))
	}

	t.Logf("err: %v", scanner.Err())
}

func BenchmarkScanner_Scan(b *testing.B) {
	file, err := os.Open("/var/log/system.log")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			scanner.Text()
		}
	}

}

func BenchmarkForwardScanner_Scan(b *testing.B) {
	file, err := os.Open("/var/log/system.log")
	if err != nil {
		b.Fatal(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		scanner := NewForwardScanner(file, fileInfo.Size())
		for scanner.Scan() {
			scanner.Text()
		}
	}
}
