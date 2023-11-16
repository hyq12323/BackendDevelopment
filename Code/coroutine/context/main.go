package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 新的需求，可以主动退出监控程序

func cpuInfo(ctx context.Context) {
	// 这里能拿到一个请求的id
	fmt.Printf("traceId: %s\r\n", ctx.Value("traceId"))
	// 记录一些日志，这次请求是哪个traceId打印的
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出cpu监控")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("cpu的信息")
		}
	}
}

func main() {
	// 渐进式的方式

	// 有一个goroutine监控cpu的信息
	wg.Add(1)
	// context包提供了三种函数，WithCancel，WithTimeout，WithValue
	// 如果你的goroutine，函数中，如果希望被控制、超时、传值，但是我不希望影响我原来的接口信息的时候，函数参数中第一个就尽量的加上一个ctx
	// 1. 主动取消，上级取消，下级全部取消
	//ctx1, cancel := context.WithCancel(context.Background())
	//ctx2, _ := context.WithCancel(ctx1)

	// 2. timeout 主动超时
	ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)

	// 3. WithDeadline 在时间点cancel

	// 4. WithValue 主动传值
	valueCtx := context.WithValue(ctx, "traceId", "gjw123")

	go cpuInfo(valueCtx)
	// cancel()
	wg.Wait()
	fmt.Println("监控完成")
}
