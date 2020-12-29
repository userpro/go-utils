/*
 * @Author: dongzhzheng
 * @Date: 2020-08-27 12:06:06
 * @LastEditTime: 2020-09-22 15:08:03
 * @LastEditors: dongzhzheng
 * @FilePath: /recommend_item_embedding/common/slice.go
 * @Description: slice相关函数
 */

package common

import "reflect"

// SliceSize interface装的slice时 获取slice的长度
func SliceSize(a interface{}) int {
	if reflect.TypeOf(a).Kind() != reflect.Slice {
		return -1
	}
	ins := reflect.ValueOf(a)
	return ins.Len()
}
