package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	//name, age, sex := user.User()

	//fmt.Printf("姓名: %s,age: %d,性别:%s", name, age, sex)

	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64

	fmt.Printf("%T,%dB,%v~%v\n", i8, unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T,%dB,%v~%v\n", i16, unsafe.Sizeof(i16), math.MinInt16, math.MaxInt16)
	fmt.Printf("%T,%dB,%v~%v\n", i32, unsafe.Sizeof(i32), math.MinInt32, math.MaxInt32)
	fmt.Printf("%T,%dB,%v~%v\n", i64, unsafe.Sizeof(i64), math.MinInt64, math.MaxInt64)

	fmt.Printf("%T,%dB,%v~%v\n", ui8, unsafe.Sizeof(ui8), 0, math.MaxUint8)
	fmt.Printf("%T,%dB,%v~%v\n", ui16, unsafe.Sizeof(ui16), 0, math.MaxUint16)
	fmt.Printf("%T,%dB,%v~%v\n", ui32, unsafe.Sizeof(ui32), 0, math.MaxUint32)
	fmt.Printf("%T,%dB,%v~%v\n", ui64, unsafe.Sizeof(ui64), 0, uint64(math.MaxUint64))

}
