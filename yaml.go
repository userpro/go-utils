/*
 * @Author: dongzhzheng
 * @Date: 2020-08-24 14:51:00
 * @LastEditTime: 2020-09-22 15:15:43
 * @LastEditors: dongzhzheng
 * @FilePath: /recommend_item_embedding/common/yaml.go
 * @Description:
 */

package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v3"
)

// YAMLDate 能从yaml解析出Date
type YAMLDate struct {
	Time time.Time
}

// UnmarshalYAML 实现接口
func (t *YAMLDate) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var buf string
	err := unmarshal(&buf)
	if err != nil {
		return nil
	}

	tt, err := time.Parse("2006-01-02", strings.TrimSpace(buf))
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

// MarshalYAML 实现接口
func (t YAMLDate) MarshalYAML() (interface{}, error) {
	return t.Time.Format("2006-01-02"), nil
}

// ToString 返回字符串
func (t YAMLDate) ToString() string {
	return t.Time.Format("2006-01-02")
}

// Set 设置
func (t *YAMLDate) Set(tm time.Time) {
	t.Time = tm
}

// GetYAML 从yaml配置文件解析到struct
func GetYAML(f string, s interface{}) error {
	ymlFile, err := os.Open(f)
	if err != nil {
		return fmt.Errorf("open yaml file(%s) failed: %v", f, err)
	}
	defer ymlFile.Close()

	ymlContent, err := ioutil.ReadAll(ymlFile)
	if err != nil {
		return fmt.Errorf("read yaml file(%s) failed: %v", f, err)
	}

	return yaml.Unmarshal(ymlContent, s)
}
