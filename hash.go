/*
 * @Author: dongzhzheng
 * @Date: 2020-08-31 11:29:57
 * @LastEditTime: 2020-09-22 14:56:28
 * @LastEditors: dongzhzheng
 * @FilePath: /recommend_item_embedding/common/hash.go
 * @Description: hash函数
 */

package common

import (
	"fmt"
	"hash/fnv"
)

// GetHash64 得到64bit hash值
func GetHash64(input []byte) (uint64, error) {
	res := fnv.New64()
	_, err := res.Write([]byte(input))
	if err != nil {
		return 0, fmt.Errorf("hash64 err(%v)", err)
	}
	return res.Sum64(), nil
}
