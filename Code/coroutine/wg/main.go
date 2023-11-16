package main

import (
	"fmt"
	"sync"
)

// 子goroutine如何通知到住的goroutine，主的goroutine如何知道子的goroutine已经结束了
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		// 需要监控多少个goroutine执行结束
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			defer wg.Done()
		}(i)
	}
	// 等到
	wg.Wait()
	fmt.Println("all done")

	//waitgroup主要用于goroutine的执行等待，Add方法要和Done方法配套
}
