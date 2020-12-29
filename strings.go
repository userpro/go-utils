/*
 * @Author: dongzhzheng
 * @Date: 2020-12-29 20:40:32
 * @LastEditTime: 2020-12-29 20:40:51
 * @LastEditors: dongzhzheng
 * @FilePath: /go-utils/strings.go
 * @Description: 字符串相关
 */

package utils

import "unsafe"

func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	b := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}

func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
