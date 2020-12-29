/*
 * @Author: dongzhzheng
 * @Date: 2020-09-01 14:58:52
 * @LastEditTime: 2020-09-22 15:14:17
 * @LastEditors: dongzhzheng
 * @FilePath: /recommend_item_embedding/common/typecast.go
 * @Description:
 */

package utils

import (
	"encoding/binary"
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

// Uint64CastString uint64转string
func Uint64CastString(a uint64) string {
	b := Uint64ToBytes(a)
	return *(*string)(unsafe.Pointer(&b))
}

// BytesToUint64 byte数组转uint64
func BytesToUint64(buf []byte) uint64 {
	return binary.BigEndian.Uint64(buf)
}

// StringCastUint64 string强转uint64
func StringCastUint64(a string) uint64 {
	b := []byte(a)
	return BytesToUint64(b)
}
