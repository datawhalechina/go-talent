/*
 * @Author: 光城
 * @Date: 2020-11-22 11:04:05
 * @LastEditors: 光城
 * @LastEditTime: 2020-11-22 11:33:18
 * @Description:
 * @FilePath: /go-talent/code/concurrency/channel/select/select_deadlock.go
 */
package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	defer close(ch)
	select {
	case <-ch:
		fmt.Println("ch 1")
	case <-ch:
		fmt.Println("ch 2")
	default:
		fmt.Println("ch default")
	}
}
