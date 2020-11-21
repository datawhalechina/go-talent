/*
 * @Author: 光城
 * @Date: 2020-11-21 15:33:44
 * @LastEditors: 光城
 * @LastEditTime: 2020-11-21 16:01:04
 * @Description:
 * @FilePath: /go-talent/code/concurrency/waitgroup/waitgroup.go
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

// 这是我们将在每个goroutine中运行的函数。
// 注意，等待组必须通过指针传递给函数。
func worker(id int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}
