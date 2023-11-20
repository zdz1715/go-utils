package goutils

import (
	"testing"
)

var testStr = `
《三国演义》是中国古典四大名著之一，是中国第一部长篇章回体历史演义小说，全名为《三国志通俗演义》，作者是元末明初的小说家罗贯中。
《三国演义》描写了从东汉末年到西晋初年之间近105年的历史风云，以描写战争为主，反映了东汉末年的群雄割据混战和魏、蜀、吴三国之间的政治和军事斗争。
反映了三国时代各类社会斗争与矛盾的转化，并概括了这一时代的历史巨变，塑造了一批叱咤风云的三国英雄人物
`
var testByte = []byte(testStr)

func BenchmarkBytesToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BytesToString(testByte)
	}
}

func BenchmarkBytesToStringByString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(testByte)
	}
}

func BenchmarkStringToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToBytes(testStr)
	}
}

func BenchmarkStringToBytesByByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []byte(testStr)
	}
}

func TestBytesToString(t *testing.T) {
	bs := StringToBytes(testStr)
	t.Logf("StringToBytes() len: %d, cap: %d", len(bs), cap(bs))
	t.Logf("[]byte() len: %d, cap: %d", len(testByte), cap(testByte))
}
