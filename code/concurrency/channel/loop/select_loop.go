/*
 * @Author: 光城
 * @Date: 2020-11-22 11:26:44
 * @LastEditors: 光城
 * @LastEditTime: 2020-11-22 12:05:17
 * @Description:
 * @FilePath: /go-talent/code/concurrency/channel/loop/select_loop.go
 */

package main

import (
	"fmt"
	"time"
)

func f1(c chan int, s chan string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		c <- i
	}
	s <- "stop"
}

func f2(c chan int, s chan string) {
	for i := 20; i >= 0; i-- {
		time.Sleep(time.Second)
		c <- i
	}
	s <- "stop"
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	signal := make(chan string, 10)

	go f1(c1, signal)
	go f2(c2, signal)
LOOP:
	for {
		select {
		case data := <-c1:
			fmt.Println("c1 data is ", data)
		case data := <-c2:
			fmt.Println("c2 data is ", data)
		case data := <-signal:
			fmt.Println("signal is ", data)
			break LOOP
		}
	}
}
