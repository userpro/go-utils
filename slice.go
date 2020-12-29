/*
 * @Author: dongzhzheng
 * @Date: 2020-08-27 12:06:06
 * @LastEditTime: 2020-12-29 20:33:22
 * @LastEditors: dongzhzheng
 * @FilePath: /go-utils/slice.go
 * @Description: slice相关函数
 */

package utils

import "reflect"

// SliceSize interface装的slice时 获取slice的长度
func SliceSize(a interface{}) int {
	if reflect.TypeOf(a).Kind() != reflect.Slice {
		return -1
	}
	ins := reflect.ValueOf(a)
	return ins.Len()
}
