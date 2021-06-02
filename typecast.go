/*
 * @Author: dongzhzheng
 * @Date: 2020-09-01 14:58:52
 * @LastEditTime: 2021-06-02 11:28:07
 * @LastEditors: dongzhzheng
 * @FilePath: /go-utils/typecast.go
 * @Description:
 */

package utils

import (
	"encoding/binary"
	"math"
	"unsafe"
)

// Int8sToBytes int8数组转byte数组
func Int8sToBytes(i []int8) []byte {
	r := []byte{}
	for _, v := range i {
		r = append(r, byte(v))
	}
	return r
}

// BytesToInt8s byte数组转int8数组
func BytesToInt8s(i []byte) []int8 {
	r := []int8{}
	for _, v := range i {
		r = append(r, int8(v))
	}
	return r
}

// Uint64ToBytes unit64转byte数组
func Uint64ToBytes(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, i)
	return buf
}

// Uint64ToString uint64转string
func Uint64ToString(a uint64) string {
	b := Uint64ToBytes(a)
	return *(*string)(unsafe.Pointer(&b))
}

// BytesToUint64 byte数组转uint64
func BytesToUint64(buf []byte) uint64 {
	return binary.BigEndian.Uint64(buf)
}

// StringToUint64 string转uint64
func StringToUint64(a string) uint64 {
	b := []byte(a)
	return BytesToUint64(b)
}

// BytesToFloat64 bytes转float64
func BytesToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

// Float64ToBytes float64转bytes
func Float64ToBytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

// BytesToFloat32 bytes转float32
func BytesToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}

// Float32ToBytes float32转bytes
func Float32ToBytes(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}
