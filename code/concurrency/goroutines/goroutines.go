/*
 * @Author: 光城
 * @Date: 2020-11-21 15:26:24
 * @LastEditors: 光城
 * @LastEditTime: 2020-11-21 15:27:33
 * @Description:
 * @FilePath: /go-talent/code/concurrency/goroutines/main.go
 */

package main

import "fmt"

func main() {
	go func() {
		fmt.Println("you forgot me !")
	}()
}
