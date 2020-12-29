/*
 * @Author: dongzhzheng
 * @Date: 2020-09-10 16:24:35
 * @LastEditTime: 2020-12-22 14:42:51
 * @LastEditors: dongzhzheng
 * @FilePath: /recommend_item_embedding/common/stack.go
 * @Description: 打印堆栈信息
 */

package common

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

// SetupSigusr1Trap 捕获到指定signal时打印出goroutine堆栈
func SetupSigusr1Trap() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go func() {
		for range c {
			log.Println(DumpStacks(16384))
		}
	}()
}

// DumpStacks dump堆栈信息
func DumpStacks(stackSize int) string {
	// buf := make([]byte, 16384)
	buf := make([]byte, stackSize)
	buf = buf[:runtime.Stack(buf, true)]
	return fmt.Sprintf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
}
