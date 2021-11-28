package studybase

import (
	"fmt"
	"unsafe"
)

func Bytes2String(bytes []byte) string {
	var b byte
	b = 101
	var a uint8
	a = b
	fmt.Println(a)

	return *(*string)(unsafe.Pointer(&bytes))
}

func String2Bytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

