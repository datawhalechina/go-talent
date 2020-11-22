/*
 * @Author: 光城
 * @Date: 2020-11-22 11:09:29
 * @LastEditors: 光城
 * @LastEditTime: 2020-11-22 11:28:44
 * @Description:
 * @FilePath: /go-talent/code/concurrency/channel/select/select_timeout.go
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool, 1)
	defer close(timeout)
	go func() {
		time.Sleep(2 * time.Second)
		timeout <- true
	}()
	ch := make(chan int, 1)

	defer close(ch)

	select {
	case <-ch:
		fmt.Println("ch 1")
	case <-timeout:
		fmt.Println("timeout 1")
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 2")
	}
}
