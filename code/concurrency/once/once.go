/*
 * @Author: 光城
 * @Date: 2020-11-21 16:22:23
 * @LastEditors: 光城
 * @LastEditTime: 2020-11-21 16:52:17
 * @Description:
 * @FilePath: /go-talent/code/once/once.go
 */

package main

import (
	"fmt"
	"sync"
)

var doOnce sync.Once

func main() {
	DoSomething()
	DoSomething()
}

func DoSomething() {
	doOnce.Do(func() {
		fmt.Println("Run once - first time, loading...")
	})
	fmt.Println("Run this every time")
}
