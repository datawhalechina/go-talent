/*
 * @Author: 光城
 * @Date: 2020-11-22 10:56:53
 * @LastEditors: 光城
 * @LastEditTime: 2020-11-22 11:28:33
 * @Description:
 * @FilePath: /go-talent/code/concurrency/channel/select/select_random.go
 */

package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	defer close(ch)
	ch <- 1
	select {
	case <-ch:
		fmt.Println("ch 1")
	case <-ch:
		fmt.Println("ch 2")
	default:
		fmt.Println("ch default")
	}
}
