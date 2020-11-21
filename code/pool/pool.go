/*
 * @Author: 光城
 * @Date: 2020-11-21 17:42:57
 * @LastEditors: 光城
 * @LastEditTime: 2020-11-21 18:01:07
 * @Description:
 * @FilePath: /go-talent/code/pool/pool.go
 */
package main

import (
	"fmt"
	"sync"
)

var pool *sync.Pool

type Foo struct {
	Name string
}

func Init() {
	pool = &sync.Pool{
		New: func() interface{} {
			return new(Foo)
		},
	}
}

func main() {
	fmt.Println("Init p")
	Init()

	p := pool.Get().(*Foo)
	fmt.Println("第一次取：", p)
	p.Name = "bob"
	pool.Put(p)

	fmt.Println("池子有对象了，调用获取", pool.Get().(*Foo))
	fmt.Println("池子空了", pool.Get().(*Foo))
}
