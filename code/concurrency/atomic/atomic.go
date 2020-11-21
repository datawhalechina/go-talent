/*
 * @Author: 光城
 * @Date: 2020-11-21 17:12:47
 * @LastEditors: 光城
 * @LastEditTime: 2020-11-21 17:41:59
 * @Description:
 * @FilePath: /go-talent/code/concurrency/atomic/atomic.go
 */

package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	// 加减
	var i int64
	atomic.AddInt64(&i, 1)
	fmt.Println("i = i + 1 =", i)
	atomic.AddInt64(&i, -1)
	fmt.Println("i = i - 1 =", i)

	// 比较并交换
	var a int32 = 1
	var b int32 = 2
	var c int32 = 3
	ok := atomic.CompareAndSwapInt32(&a, a, b)
	fmt.Printf("ok = %v, a = %v, b = %v\n", ok, a, b)
	ok = atomic.CompareAndSwapInt32(&a, c, b)
	fmt.Printf("ok = %v, a = %v, b = %v, c = %v\n", ok, a, b, c)

	// 交换
	var x int32 = 1
	var y int32 = 2
	old := atomic.SwapInt32(&x, y)
	fmt.Println(x, old)

	// 加载
	var x1 int32 = 1
	y1 := atomic.LoadInt32(&x1)
	fmt.Println("x1, y1:", x1, y1)

	// 存储
	var xx int32 = 1
	var yy int32 = 2
	atomic.StoreInt32(&yy, atomic.LoadInt32(&xx))
	fmt.Println(xx, yy)

	// 原子类型
	v := atomic.Value{}
	v.Store(1)
	fmt.Println(v.Load())
}
