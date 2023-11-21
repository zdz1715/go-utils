package fileutils

import (
	"fmt"
	"testing"
)

func TestFileLineFunc(t *testing.T) {
	//err := ReadFileLineFromEndFunc("/tmp/clbh64ko47mieriluks0", func(file *os.File, line string) {
	//	fmt.Println(line)
	//})
	//if err != nil {
	//	t.Fatal(err)
	//}

	s := make([]int, 2, 2)
	s = append(s, 2)
	fmt.Println(s[0], s[1])
}

func TestIsDirE(t *testing.T) {
	if err := IsDirE("./file.go"); err != nil {
		t.Fatal(err)
	}
}
