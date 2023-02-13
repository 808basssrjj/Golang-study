package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	v1 := uint(32)
	v2 := int(13)
	fmt.Println(reflect.TypeOf(v1))
	fmt.Println(reflect.TypeOf(v2))
	fmt.Println(reflect.TypeOf(&v1))
	fmt.Println(reflect.TypeOf(&v2))

	// 1.把*int转为*uint
	// unsafe.Pointer可以和各种指针类型相互转换，也可以转换为uintptr类型
	p := &v1
	p = (*uint)(unsafe.Pointer(&v2))
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(*p)

	// 2.string和[]byte零拷贝转换
	s := "Hello world";
	b := stringToBytes(s)
	fmt.Println(b)
	s = BytesToString(b)
	fmt.Println(s)

	// 1.20版本 reflect.SliceHeader 和 reflect.StringHeader 将会被标注为被废弃
	// func StringToBytes(s string) []byte {
	//     return unsafe.Slice(unsafe.StringData(s), len(s))
	// }

	// func BytesToString(b []byte) string {
	//     return unsafe.String(&b[0], len(b))
	// }
}


func stringToBytes(s string) []byte {
	stringH := (*reflect.StringHeader)(unsafe.Pointer(&s))

	bh := reflect.SliceHeader{
		Data: stringH.Data,
		Len:  stringH.Len,
		Cap:  stringH.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&bh))
}

func BytesToString(b []byte) string {
	sliceH := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	sh := reflect.StringHeader{
		Data: sliceH.Data,
		Len: sliceH.Len,
	}

	return *(*string)(unsafe.Pointer(&sh))
}
