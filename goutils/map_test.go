package goutils

import (
	"strconv"
	"testing"
)

func BenchmarkKeys(b *testing.B) {
	in := make(map[string]int)
	for i := 1; i <= 10; i++ {
		in[strconv.Itoa(i)] = i
	}
	for i := 0; i < b.N; i++ {
		Keys(in)
	}
}

func BenchmarkValues(b *testing.B) {
	in := make(map[string]int)
	for i := 1; i <= 10; i++ {
		in[strconv.Itoa(i)] = i
	}
	for i := 0; i < b.N; i++ {
		Values(in)
	}
}
