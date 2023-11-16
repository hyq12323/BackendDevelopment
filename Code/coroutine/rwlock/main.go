package main

import (
	"fmt"
	"sync"
	"time"
)

// 锁本质是将并行的代码串行化，使用lock肯定回影响性能
// 即使是设计锁，那么也应该尽量保证并行
// 我们有两组协程，其中一组复制写数据，另一组复制读数据，web系统中绝大部分场景都是读多写少
// 虽然有多个goroutine，但是仔细分析我们会发现，读协程之间应该并发，读和写应该串行，写和写之间也不应该并行
// 读写锁

func main() {
	// var num int
	var rwLock sync.RWMutex
	var wg sync.WaitGroup

	wg.Add(6)
	// 写的goroutine
	go func() {
		time.Sleep(time.Second * 3)
		defer wg.Done()
		rwLock.Lock() // 加写锁，写锁会防止别的写锁获取，和读写获取
		defer rwLock.Unlock()
		time.Sleep(time.Second * 5)
		//num = 12
		fmt.Println("get write lock")
	}()

	// 读的goroutine
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for {
				rwLock.RLock()
				time.Sleep(500 * time.Millisecond)
				fmt.Println("get read lock")
				rwLock.RUnlock() // 加读锁，读锁不会阻止别人的的读
			}
			//fmt.Println(num)
		}()
	}
	wg.Wait()
}
