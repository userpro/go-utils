/*
 * @Author: dongzhzheng
 * @Date: 2020-08-26 20:07:15
 * @LastEditTime: 2020-09-22 15:09:49
 * @LastEditors: dongzhzheng
 * @FilePath: /recommend_item_embedding/common/time.go
 * @Description:
 */

package common

import "time"

const (
	GOTIMEFORMAT = "2006-01-02 15:04:05"
)

// TimeToStr time.Time => string
func TimeToStr(t time.Time) string {
	return t.Format(GOTIMEFORMAT)
}

// StrToTime string => time.Time
func StrToTime(t string) time.Time {
	loc, _ := time.LoadLocation("Local")
	t1, _ := time.ParseInLocation(GOTIMEFORMAT, t, loc) //使用模板在对应时区转化为time.time类型
	return t1
}
