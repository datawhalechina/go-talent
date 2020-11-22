/*
 * @Author: 光城
 * @Date: 2020-11-22 11:17:53
 * @LastEditors: 光城
 * @LastEditTime: 2020-11-22 11:29:07
 * @Description:
 * @FilePath: /go-talent/code/concurrency/channel/check/select_block.go
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
	case ch <- 1:
		fmt.Println("channel value is ", <-ch)
		fmt.Println("channel value is ", <-ch)
	default:
		fmt.Println("channel blocking")
	}
}
