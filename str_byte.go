package util

import "unsafe"

// string转[]byte无拷贝
func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// Bytes2Str byte数组直接转成string对象，不发生内存copy
func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Uint82Str uint8数组直接转成string对象,不发生内存copy
func Uint82Str(u []uint8) string {
	return *(*string)(unsafe.Pointer(&u))
}

func ByteToInt8(s []byte) []int8 {
	d := *(*[]int8)(unsafe.Pointer(&s))
	return d
}
func Int8ToByte(s []int8) []byte {
	d := *(*[]byte)(unsafe.Pointer(&s))
	return d
}
func StringToByte(s string) []byte {
	return []byte(s)
}
func ByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
func StringToInt8(s string) []int8 {
	var b []byte
	b = []byte(s)
	return *(*[]int8)(unsafe.Pointer(&b))
}
func Int8ToString(b []int8) string {
	return ByteToString(Int8ToByte(b))
}
