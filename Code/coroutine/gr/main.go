package main

import (
	"fmt"
	"time"
)

// 主协程
func main() {
	// 主死随从

	// 匿名函数启动goroutine

	// 1. 闭包
	// 2. for循环的问题,for循环的时候每个变量会重用
	// 每次for循环的时候，i变量会被重用，当进行到第二轮的for循环的时候，这个i就变了
	for i := 0; i < 100; i++ {
		// 第一种解决办法 tmp := i
		go func(i int) {
			fmt.Println(i)
		}(i) //值传递
	}

	fmt.Println("main goroutine")
	time.Sleep(10 * time.Second)
}
