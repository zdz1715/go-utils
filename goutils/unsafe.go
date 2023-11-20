package goutils

import (
	"unsafe"
)

// BytesToString converts byte slice to string.
func BytesToString(s []byte) string {
	//return *(*string)(unsafe.Pointer(&b))
	// go 1.20+
	return unsafe.String(unsafe.SliceData(s), len(s))
}

// StringToBytes converts string to byte slice.
func StringToBytes(s string) []byte {
	//return *(*[]byte)(unsafe.Pointer(
	//	&struct {
	//		string
	//		Cap int
	//	}{s, len(s)},
	//))

	// go 1.20+
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
