/*
 * @Author: 光城
 * @Date: 2020-11-21 16:01:29
 * @LastEditors: 光城
 * @LastEditTime: 2020-11-21 16:13:39
 * @Description:
 * @FilePath: /go-talent/code/concurrency/mutex/mutex.go
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter 的并发使用是安全的。
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc 增加给定 key 的计数器的值。
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.v[key]++
}

// Value 返回给定 key 的计数器的当前值。
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
